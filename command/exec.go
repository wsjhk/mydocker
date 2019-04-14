package command

import (
	"fmt"
	"os"
	"os/exec"
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
	//nsenter.EnterNamespace()
	cmd := exec.Command("/proc/self/exe", "exec")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Errorf("Exec container %s error %v", containerName, err)
	}
}