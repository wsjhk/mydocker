package command

import (
	"fmt"
	"github.com/nicktming/mydocker/cgroups"
	"log"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func Run(command string, tty bool, cg *cgroups.CroupManger, rootPath string)  {
	//cmd := exec.Command(command)

	reader, writer, err := os.Pipe()
	if err != nil {
		log.Printf("Error: os.pipe() error:%v\n", err)
		return
	}

	//cmd := exec.Command("/proc/self/exe", "init", command)

	cmd := exec.Command("/proc/self/exe", "init")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	newRootPath := getRootPath(rootPath)
	cmd.Dir = newRootPath + "/busybox"
	if err := NewWorkDir(newRootPath); err != nil {
		cmd.Dir = newRootPath + "/mnt"
	}

	cmd.ExtraFiles = []*os.File{reader}
	sendInitCommand(command, writer)

	if tty {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
	}
	/**
	 *   Start() will not block, so it needs to use Wait()
	 *   Run() will block
	 */
	if err := cmd.Start(); err != nil {
		log.Printf("Run Start err: %v.\n", err)
		log.Fatal(err)
	}
	//log.Printf("222 before process pid:%d, memory:%s\n", cmd.Process.Pid, memory)

	//subsystems.Set(memory)
	//subsystems.Apply(strconv.Itoa(cmd.Process.Pid))
	//defer subsystems.Remove()

//	sendInitCommand(command, writer)

//	cmd.Dir = "/root"

	cg.Set()
	defer cg.Destroy()
	cg.Apply(strconv.Itoa(cmd.Process.Pid))

	cmd.Wait()
}

func sendInitCommand(command string, writer *os.File)  {
	_, err := writer.Write([]byte(command))
	if err != nil {
		log.Printf("writer.Write error:%v\n", err)
		return
	}
	writer.Close()
}


func getRootPath(rootPath string) string {
	log.Printf("rootPath:%s\n", rootPath)
	defaultPath := "/root"
	if rootPath == "" {
		log.Printf("rootPath is empaty, set cmd.Dir by default: /root/busybox\n")
		return defaultPath
	}
	imageTar := rootPath + "/busybox.tar"
	exist, _ := PathExists(imageTar)
	if !exist {
		log.Printf("%s does not exist, set cmd.Dir by default: /root/busybox\n", imageTar)
		return defaultPath
	}
	imagePath := rootPath + "/busybox"
	exist, _ = PathExists(imageTar)
	if exist {
		os.RemoveAll(imagePath)
	}
	if err := os.Mkdir(imagePath, 0777); err != nil {
		log.Printf("mkdir %s err:%v, set cmd.Dir by default: /root/busybox\n", imagePath, err)
		return defaultPath
	}
	if _, err := exec.Command("tar", "-xvf", imageTar, "-C", imagePath).CombinedOutput(); err != nil {
		log.Printf("tar -xvf %s -C %s, err:%v, set cmd.Dir by default: /root/busybox\n", imageTar, imagePath, err)
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

func NewWorkDir(rootPath string) error {
	if err := CreateContainerLayer(rootPath); err != nil {
		return fmt.Errorf("CreateContainerLayer(%s) error: %v.\n", rootPath, err)
	}
	if err := CreateMntPoint(rootPath); err != nil {
		return fmt.Errorf("CreateContainerLayer(%s) error: %v.\n", rootPath, err)
	}
	if err := SetMountPoint(rootPath); err != nil {
		return fmt.Errorf("CreateContainerLayer(%s) error: %v.\n", rootPath, err)
	}
	return nil
}

func CreateContainerLayer(rootPath string) error {
	writerLayer := rootPath + "/writerLayer"
	if err := os.Mkdir(writerLayer, 0777); err != nil {
		log.Printf("mkdir %s err:%v\n", writerLayer, err)
		return fmt.Errorf("mkdir %s err:%v\n", writerLayer, err)
	}
	return nil 
}

func CreateMntPoint(rootPath string) error {
	mnt := rootPath + "/mnt"
	if err := os.Mkdir(mnt, 0777); err != nil {
		log.Printf("mkdir %s err:%v\n", mnt, err)
		return fmt.Errorf("mkdir %s err:%v\n", mnt, err)
	}
	return nil
}

func SetMountPoint(rootPath string) error {
	dirs := "dirs=" + rootPath + "/writerLayer:" + rootPath + "/busybox"
	mnt := rootPath + "/mnt"
	if _, err := exec.Command("mount", "-t", "aufs", "-o", dirs, "none", mnt).CombinedOutput(); err != nil {
		log.Printf("mount -t aufs -o %s none %s, err:%v\n", dirs, mnt, err)
		return fmt.Errorf("mount -t aufs -o %s none %s, err:%v\n", dirs, mnt, err)
	}
	return nil
}