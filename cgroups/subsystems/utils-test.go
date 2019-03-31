package subsystems

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

func Test000(t *testing.T)  {
	mountPath := FindCgroupMountPoint("memory")
	log.Printf("mountPath:%s\n", mountPath)
}

func Test001(t *testing.T)  {
	absolutePath := FindAbsolutePath("memory")
	log.Printf("absolutePath:%s\n", absolutePath)
}

func Test002(t *testing.T)  {
	Set("10M")
	pid := os.Getpid()
	log.Printf("current pid : %s\n", strconv.Itoa(pid))
	Apply(strconv.Itoa(pid))
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
	}
}