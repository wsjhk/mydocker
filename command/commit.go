package command

import (
	"log"
	"os/exec"
)

func Commit(imageName string) {
	mntPath := DEFAULTPATH + "/mnt"
	imageTar := DEFAULTPATH + "/" + imageName + ".tar"
	log.Printf("imageTar:%s\n", imageTar)
	if _, err := exec.Command("tar", "-czf", imageTar, "-C", mntPath, ".").CombinedOutput(); err != nil {
		log.Printf("Error: tar -czf %s -C %s .; err:%v\n", imageTar, mntPath, err)
	}
}
