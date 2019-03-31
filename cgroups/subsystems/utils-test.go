package subsystems

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

func Test001(t *testing.T)  {
	//paths := FindAbsolutePath("memory")
	//log.Printf("paths:%s\n", paths)
	Set("10M")
	pid := os.Getpid()
	log.Printf("current pid : %s\n", strconv.Itoa(pid))
	Apply(strconv.Itoa(pid))
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
	}
}