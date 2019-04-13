package command

import (
	"fmt"
	"os"
)

var (
	MYDOCKER_PID = "mydocker_pid"
	MYDOCKER_COMMAND = "mydocker_cmd"
)

func Exec(containerName, command string) {
	containerInfo, err := GetContainerInfo(containerName)
	if err != nil {
		fmt.Errorf("GetContainerInfo error:%v\n", err)
		return
	}
	pid := containerInfo.Pid
	os.Setenv(MYDOCKER_PID, pid)
	os.Setenv(MYDOCKER_COMMAND, command)
}