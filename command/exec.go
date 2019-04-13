package command

import (
	"fmt"
	"github.com/nicktming/mydocker/nsenter"
	"os"
)

func Exec(containerName, command string) {
	containerInfo, err := GetContainerInfo(containerName)
	if err != nil {
		fmt.Errorf("GetContainerInfo error:%v\n", err)
		return
	}
	pid := containerInfo.Pid
	os.Setenv("mydocker_pid", pid)
	os.Setenv("mydocker_cmd", command)
	nsenter.EnterNamespace()
}