package command

import (
	"fmt"
	"github.com/nicktming/mydocker/cgroups"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

const (
	DEFAULTPATH = "/nicktming"
)

func Run(command string, tty bool, cg *cgroups.CroupManger, rootPath string, volumes []string, containerName, imageName string)  {
	//cmd := exec.Command(command)

	reader, writer, err := os.Pipe()
	if err != nil {
		log.Printf("Error: os.pipe() error:%v\n", err)
		return
	}

	//cmd := exec.Command("/proc/self/exe", "init", command)

	initCmd, err := os.Readlink("/proc/self/exe")
	if err != nil {
		fmt.Errorf("get init process error %v", err)
		return
	}

	cmd := exec.Command(initCmd, "init")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}




	cmd.ExtraFiles = []*os.File{reader}
	sendInitCommand(command, writer)

	id := ContainerUUID()
	if containerName == "" {
		containerName = id
	}

	//log.Printf("volume:%s\n", volumes)

	newRootPath := getRootPath(rootPath, imageName)
	//cmd.Dir = newRootPath + "/busybox"
	if err := NewWorkDir(newRootPath, containerName, imageName, volumes); err == nil {
		cmd.Dir = newRootPath + "/mnt/" + containerName
	} else {
		log.Printf("NewWorkDir error:%v\n", err)
		return
	}
	defer ClearWorkDir(newRootPath, containerName, volumes)

	if tty {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
	} else {
		logFile, err := GetLogFile(containerName)
		if err != nil {
			fmt.Errorf("GetLogFile error:%v\n", err)
			return
		}
		cmd.Stdout = logFile
	}
	/**
	 *   Start() will not block, so it needs to use Wait()
	 *   Run() will block
	 */
	if err := cmd.Start(); err != nil {
		log.Printf("Run Start err: %v.\n", err)
		log.Fatal(err)
	}

	cg.Set()
	defer cg.Destroy()
	cg.Apply(strconv.Itoa(cmd.Process.Pid))

	RecordContainerInfo(strconv.Itoa(cmd.Process.Pid), containerName, id, command)

	// false 表明父进程(Run程序)无须等待子进程(Init程序,Init进程后续会被用户程序覆盖)
	if tty {
		cmd.Wait()
		DeleteContainerInfo(containerName)
	}
}

func sendInitCommand(command string, writer *os.File)  {
	_, err := writer.Write([]byte(command))
	if err != nil {
		log.Printf("writer.Write error:%v\n", err)
		return
	}
	writer.Close()
}


func getRootPath(rootPath, imageName string) string {
	//log.Printf("rootPath:%s\n", rootPath)
	defaultPath := DEFAULTPATH
	if rootPath == "" {
		log.Printf("rootPath is empaty, set rootPath: %s\n", defaultPath)
		rootPath = defaultPath
	}
	imageTar := rootPath + "/" + imageName + ".tar"
	exist, _ := PathExists(imageTar)
	if !exist {
		log.Printf("%s does not exist, set cmd.Dir by default: %s/mnt\n", imageTar, defaultPath)
		return defaultPath
	}
	imagePath := rootPath + "/" + imageName
	if err := os.MkdirAll(imagePath, 0777); err != nil {
		log.Printf("mkdir %s err:%v, set cmd.Dir by default: %s/mnt\n", imagePath, err, defaultPath)
		return defaultPath
	}
	if _, err := exec.Command("tar", "-xvf", imageTar, "-C", imagePath).CombinedOutput(); err != nil {
		log.Printf("tar -xvf %s -C %s, err:%v, set cmd.Dir by default: %s/mnt\n", imageTar, imagePath, err, defaultPath)
		return defaultPath
	}
	return rootPath
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ClearWorkDir(rootPath, containerName string, volumes []string)  {
	for _, volume := range volumes {
		ClearVolume(rootPath, volume, containerName)
	}
	ClearMountPoint(rootPath, containerName)
	ClearWriterLayer(rootPath, containerName)
}

func ClearVolume(rootPath, volume, containerName string)  {
	if volume != "" {
		containerMntPath   := rootPath + "/mnt/" + containerName
		mountPath 	  := strings.Split(volume, ":")[1]
		containerPath := containerMntPath + mountPath
		if _, err := exec.Command("umount", "-f", containerPath).CombinedOutput(); err != nil {
			log.Printf("umount -f %s, err:%v\n", containerPath, err)
		}
		if err := os.RemoveAll(containerPath); err != nil {
			log.Printf("remove %s, err:%v\n", containerPath, err)
		}
	}
}

func ClearMountPoint(rootPath, containerName string)  {
	mnt := rootPath + "/mnt/" + containerName
	if _, err := exec.Command("umount", "-f", mnt).CombinedOutput(); err != nil {
		log.Printf("mount -f %s, err:%v\n", mnt, err)
	}
	if err := os.RemoveAll(mnt); err != nil {
		log.Printf("remove %s, err:%v\n", mnt, err)
	}
}

func ClearWriterLayer(rootPath, containerName string) {
	writerLayer := rootPath + "/writerLayer/" + containerName
	if err := os.RemoveAll(writerLayer); err != nil {
		log.Printf("remove %s, err:%v\n", writerLayer, err)
	}
}

func NewWorkDir(rootPath, containerName, imageName string, volumes []string) error {
	if err := CreateContainerLayer(rootPath, containerName); err != nil {
		return fmt.Errorf("CreateContainerLayer(%s) error: %v.\n", rootPath, err)
	}
	if err := CreateMntPoint(rootPath, containerName); err != nil {
		return fmt.Errorf("CreateMntPoint(%s) error: %v.\n", rootPath, err)
	}
	if err := SetMountPoint(rootPath, containerName, imageName); err != nil {
		return fmt.Errorf("SetMountPoint(%s) error: %v.\n", rootPath, err)
	}
	for _, volume := range volumes {
		if err := CreateVolume(rootPath, volume, containerName); err != nil {
			return fmt.Errorf("CreateVolume(%s, %s) error: %v.\n", rootPath, volume, err)
		}
	}
	return nil
}

func CreateVolume(rootPath, volume, containerName string) error {
	if volume != "" {
		containerMntPath := rootPath + "/mnt/" + containerName
		hostPath 	:= strings.Split(volume, ":")[0]
		exist, _ := PathExists(hostPath)
		if !exist {
			if err := os.MkdirAll(hostPath, 0777); err != nil {
				log.Printf("mkdir %s err:%v\n", hostPath, err)
				return fmt.Errorf("mkdir %s err:%v\n", hostPath, err)
			}
		}
		mountPath 	:= strings.Split(volume, ":")[1]
		containerPath := containerMntPath + mountPath
		if err := os.MkdirAll(containerPath, 0777); err != nil {
			log.Printf("mkdir %s err:%v\n", containerPath, err)
			return fmt.Errorf("mkdir %s err:%v\n", containerPath, err)
		}
		dirs := "dirs=" + hostPath
		if _, err := exec.Command("mount", "-t", "aufs", "-o", dirs, "none", containerPath).CombinedOutput(); err != nil {
			log.Printf("mount -t aufs -o %s none %s, err:%v\n", dirs, containerPath, err)
			return fmt.Errorf("mount -t aufs -o %s none %s, err:%v\n", dirs, containerPath, err)
		}
	}
	return nil
}

func CreateContainerLayer(rootPath, containerName string) error {
	writerLayer := rootPath + "/writerLayer/" + containerName
	if err := os.MkdirAll(writerLayer, 0777); err != nil {
		log.Printf("mkdir %s err:%v\n", writerLayer, err)
		return fmt.Errorf("mkdir %s err:%v\n", writerLayer, err)
	}
	return nil 
}

func CreateMntPoint(rootPath, containerName string) error {
	mnt := rootPath + "/mnt/" + containerName
	if err := os.MkdirAll(mnt, 0777); err != nil {
		log.Printf("mkdir %s err:%v\n", mnt, err)
		return fmt.Errorf("mkdir %s err:%v\n", mnt, err)
	}
	return nil
}

func SetMountPoint(rootPath, containerName, imageName string) error {
	dirs := "dirs=" + rootPath + "/writerLayer/" + containerName + ":" + rootPath + "/" + imageName
	mnt := rootPath + "/mnt/" + containerName
	if _, err := exec.Command("mount", "-t", "aufs", "-o", dirs, "none", mnt).CombinedOutput(); err != nil {
		log.Printf("mount -t aufs -o %s none %s, err:%v\n", dirs, mnt, err)
		return fmt.Errorf("mount -t aufs -o %s none %s, err:%v\n", dirs, mnt, err)
	}
	return nil
}