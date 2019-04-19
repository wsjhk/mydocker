package command

import (
	"log"
	"os/exec"
	"strings"
)

func Copy(source, destination string)  {
	f1 := strings.Contains(source, ":")
	f2 := strings.Contains(destination, ":")
	if (f1 && f2) || (!f1 && !f2) {
		log.Printf("f1:%v, f2:%v, not correct format\n")
		return
	}

	from_container_to_host := true
	containerUrl := source
	hostUrl 	 := destination
	if f2 {
		from_container_to_host = false
		containerUrl = destination
		hostUrl = source
	}
	containerName := strings.Split(containerUrl, ":")[0]
	containerPath := strings.Split(containerUrl, ":")[1]
	log.Printf("containerUrl:%s, hostUrl:%s, conatinerName:%s, containerPath:%s\n", containerUrl, hostUrl, containerName, containerPath)

	containerInfo, err := GetContainerInfo(containerName)
	if err != nil {
		log.Printf("GetContainerInfo error:%v\n", err)
		return
	}
	containerMntPath := containerInfo.RootPath + "/mnt" + containerName + containerPath
	hostPath 	     := hostUrl
	log.Printf("containerPath:%s, hostPath:%s\n", containerMntPath, hostPath)

	if from_container_to_host {
		FileCopy(containerMntPath, hostPath)
	} else {
		FileCopy(hostPath, containerMntPath)
	}
}

func FileCopy(src, dst string) {
	exist, _ := PathExists(src)
	if !exist {
		log.Printf("src:%s not exists!\n", src)
		return
	}
	exist, _ = PathExists(dst)
	if !exist {
		log.Printf("dst:%s not exists!\n", src)
		return
	}
	if _, err := exec.Command("cp", "-r", src, dst).CombinedOutput(); err != nil {
		log.Printf("cp -r %s %s, err:%v\n", src, dst, err)
		return
	}
}

