package cgroups

import (
	"github.com/nicktming/mydocker/cgroups/subsystems"
	"log"
	"os"
	"strconv"
	"testing"
)

// go test -v utils_test.go -test.run Test000

func Test000(t *testing.T)  {
	mountPath := subsystems.FindCgroupMountPoint("memory")
	log.Printf("mountPath:%s\n", mountPath)
}

func Test001(t *testing.T)  {
	absolutePath := subsystems.FindAbsolutePath("memory")
	log.Printf("absolutePath:%s\n", absolutePath)
}

func Test002(t *testing.T)  {
	//subsystems.Set("10M")
	//pid := os.Getpid()
	//log.Printf("current pid : %s\n", strconv.Itoa(pid))
	//subsystems.Apply(strconv.Itoa(pid))
	//for i := 0; i < 100; i++ {
	//	time.Sleep(1 * time.Second)
	//}
}

func Test003(t *testing.T)  {
	memory := "10M"
	res := subsystems.ResourceConfig{
		MemoryLimit: memory,
	}
	cg := CroupManger {
		Resource: &res,
		SubsystemsIns: make([]subsystems.Subsystem, 0),
	}
	cg.SubsystemsIns = append(cg.SubsystemsIns, &subsystems.MemorySubsystem{})

	pid := os.Getpid()
	log.Printf("current pid : %s\n", strconv.Itoa(pid))

	cg.Set()
	defer cg.Destroy()
	cg.Apply(strconv.Itoa(pid))
}