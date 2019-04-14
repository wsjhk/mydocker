package command

import "fmt"

func Remove(containerName string)  {
	containerInfo, err := GetContainerInfo(containerName)
	if err != nil {
		fmt.Errorf("GetContainerInfo error:%v\n", err)
		return
	}
	if containerInfo.Status != STOP {
		fmt.Errorf("Could not remove not stopped container!\n")
		return
	}
	RemoveContainerInfo(containerInfo)
}
