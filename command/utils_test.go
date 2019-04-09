package command

import (
	"encoding/json"
	"fmt"
	"log"
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
	RecordContainerInfo(uuid, uuid, uuid, "/bin/top")
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
	return nil
}

func ContainerUUID() string {
	str := time.Now().UnixNano()
	containerId := fmt.Sprintf("%d", str)
	log.Printf("containerId:%s\n", containerId)
	return containerId
}