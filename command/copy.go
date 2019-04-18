package command

import (
	"log"
	"strings"
)

func Copy(source, destination string)  {
	f1 := strings.Contains(source, ":")
	f2 := strings.Contains(destination, ":")
	if (f1 && f2) || (!f1 && !f2) {
		log.Printf("f1:%v, f2:%v, not correct format\n")
		return
	}

	containerUrl := source
	hostUrl 	 := destination
	if f2 {
		containerUrl = destination
		hostUrl = source
	}
	containerName := strings.Split(containerUrl, ":")[0]
	containerPath := strings.Split(containerUrl, ":")[1]
	log.Printf("containerUrl:%s, hostUrl:%s, conatinerName:%s, containerPath:%s\n", containerUrl, hostUrl, containerName, containerPath)


}

