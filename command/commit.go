package command

import (
	"fmt"
	"log"
	"os/exec"
)

func Commit(containerName, imageName string) {
	//mntPath := DEFAULTPATH + "/mnt"
	//imageTar := DEFAULTPATH + "/" + imageName + ".tar"
	containerInfo, err := GetContainerInfo(containerName)
	if err != nil {
		fmt.Errorf("GetContainerInfo error:%v\n", err)
		return
	}
	mntPath  := containerInfo.RootPath + "/mnt/" + containerName
	imageTar := containerInfo.RootPath + "/" + imageName + ".tar"
	log.Printf("mntPath:%s, imageTar:%s\n", mntPath, imageTar)
	if _, err := exec.Command("tar", "-czf", imageTar, "-C", mntPath, ".").CombinedOutput(); err != nil {
		log.Printf("Error: tar -czf %s -C %s .; err:%v\n", imageTar, mntPath, err)
	}
}
