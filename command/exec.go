package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
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

	containerEnvs := getEnvsByPid(containerInfo.Pid)
	cmd.Env = append(os.Environ(), containerEnvs...)

	if err := cmd.Run(); err != nil {
		fmt.Errorf("Exec container %s error %v", containerName, err)
	}
}

func getEnvsByPid(pid string) []string {
	path := fmt.Sprintf("/proc/%s/environ", pid)
	contentBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Read file %s error %v", path, err)
		return nil
	}
	envs := strings.Split(string(contentBytes), "\u0000")
	return envs
}
