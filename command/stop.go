package command

import (
	"log"
	"strconv"
	"syscall"
)

func Stop(containerName string)  {
	log.Printf("===>containerName:%s\n", containerName)

	containerInfo, err := GetContainerInfo(containerName)
	if err != nil {
		log.Printf("GetContainerInfo error:%v\n", err)
		return
	}
	if containerInfo.Pid == "" {
		log.Printf("container not exists!\n")
		return
	}
	pid, err := strconv.Atoi(containerInfo.Pid)
	if err != nil {
		log.Printf("strconv.Atoi(%s) error : %v\n", containerInfo.Pid, err)
		return
	}
	log.Printf("===>before kill\n")
	if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
		log.Printf("Stop container %s error %v.\n", containerName, err)
		return
	}
	log.Printf("===>after kill\n")
	containerInfo.Status = STOP
	containerInfo.Pid = ""
	UpdateContainerInfo(containerInfo)

	log.Printf("rootPath:%s\n", containerInfo.RootPath)
	log.Println(containerInfo.Volumes)
	ClearWorkDir(containerInfo.RootPath, containerName, containerInfo.Volumes)
}
