package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/rand"
	"log"
	"math"
	"os"
	"testing"
	"time"
)

type ContainerInfo struct {
	Pid			string	`json:"pid"`
	Id			string	`json:"id"`
	Name		string	`json:"name"`
	Command		string	`json:"command"`
	CreateTime	string	`json:"createTime"`
	Status		string	`json:"status"`
}

var (
	RUNNING			string = "running"
	STOP			string = "stopped"
	EXIT			string = "exited"
	INFOLOCATION	string = "/var/run/mydocker/%s"
	CONFIGNAME		string = "config.json"
)


func Test001(t *testing.T)  {
	uuid := ContainerUUID()
	if err := RecordContainerInfo(uuid, uuid, uuid, "/bin/top"); err != nil {
		log.Printf("RecordContainerInfo error : %v\n", err)
	} else {
		log.Printf("write successfully!\n")
	}
}

func RecordContainerInfo(pid, name, id, command string) error {
	containerInfo := &ContainerInfo {
		Pid: 		pid,
		Id:  		id,
		Name:		name,
		Command: 	command,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		Status:		RUNNING,
	}
	jsonInfo, _ := json.Marshal(containerInfo)
	log.Printf("jsonInfo:%s\n", string(jsonInfo))
	location := fmt.Sprintf(INFOLOCATION, name)
	file 	 := location + "/" + CONFIGNAME
	if err := os.MkdirAll(location, 0622); err != nil {
		return fmt.Errorf("create %s error : %v\n", location, err)
	}

	if err := ioutil.WriteFile(file, []byte(jsonInfo), 0622); err != nil {
		return fmt.Errorf("write %s to %s error:%v\n", jsonInfo, file, err)
	}
	return nil
}

func ContainerUUID() string {
	str := time.Now().UnixNano()
	containerId := fmt.Sprintf("%d%d", str, int(math.Abs(float64(rand.Intn(10)))))
	log.Printf("containerId:%s\n", containerId)
	return containerId
}