package command

import (
	"fmt"
	"log"
	"strconv"
	"syscall"
)

func Stop(containerName string)  {
	containerInfo, err := GetContainerInfo(containerName)
	if err != nil {
		fmt.Errorf("GetContainerInfo error:%v\n", err)
		return
	}
	if containerInfo.Pid == "" {
		log.Printf("container not exists!\n")
		return
	}
	pid, err := strconv.Atoi(containerInfo.Pid)
	if err != nil {
		fmt.Errorf("strconv.Atoi(%s) error : %v\n", containerInfo.Pid, err)
		return
	}
	if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
		fmt.Errorf("Stop container %s error %v.\n", containerName, err)
		return
	}
	containerInfo.Status = STOP
	containerInfo.Pid = ""
	UpdateContainerInfo(containerInfo)

	log.Printf("rootPath:%s\n", containerInfo.RootPath)
	log.Println(containerInfo.Volumes)
	ClearWorkDir(containerInfo.RootPath, containerName, containerInfo.Volumes)
}
