package subsystems

import (
	"bufio"
	"github.com/nicktming/mydocker/cgroups"
	"io"
	"log"
	"os"
	"strings"
)

func FindAbsolutePath(subsystem string) string {
	path := FindCgroupMountPoint(subsystem)
	if path != "" {
		absolutePath := path + "/" + cgroups.ResourceName
		exist, err := PathExists(absolutePath)
		if err != nil {
			log.Printf("PathExists error : %v\n", err)
			return ""
		}
		if !exist {
			err := os.Mkdir(absolutePath, 0755)
			if err != nil {
				log.Printf("Mkdir absolutePath:%s error : %v\n", err)
				return ""
			}
		}
		return absolutePath
	}
	return ""
}

func FindCgroupMountPoint(subsystem string) string {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		log.Printf("Error open file error : %v\n", err)
		return ""
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return ""
			}
		}
		parts := strings.Split(string(line), " ")
		if strings.Contains(parts[len(parts) - 1], subsystem) {
			return parts[4]
		}
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

