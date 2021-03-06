package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

func ShowAllContainers() {
	files, err := ioutil.ReadDir(CONTAINS)
	if err != nil {
		fmt.Errorf("readDir error : %v\n", err)
		return
	}
	var containers []*ContainerInfo
	for _, file := range files {
		container, err := GetContainerInfo(file.Name())
		if err != nil {
			log.Printf("ERROR: %v\n", err)
			continue
		}
		containers = append(containers, container)
	}
	w := tabwriter.NewWriter(os.Stdout, 12, 1, 3, ' ', 0)
	fmt.Fprint(w, "ID\tNAME\tPID\tSTATUS\tCOMMAND\tCREATED\n")
	for _, item := range containers {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
			item.Id,
			item.Name,
			item.Pid,
			item.Status,
			item.Command,
			item.CreateTime)
	}
	if err := w.Flush(); err != nil {
		fmt.Errorf("Flush error %v", err)
	}
}

func DeleteContainerInfo(name string) error {
	location := fmt.Sprintf(INFOLOCATION, name)
	if err := os.RemoveAll(location); err != nil {
		return fmt.Errorf("RemoveAll %s error:%v\n", location, err)
	}
	return nil
}

func GetContainerInfo(name string) (*ContainerInfo, error) {
	location := fmt.Sprintf(INFOLOCATION, name)
	file 	 := location + "/" + CONFIGNAME
	containerInfo := &ContainerInfo {}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("read data %s error:%s\n", data, err)
	}
	json.Unmarshal(data, containerInfo)
	return containerInfo, nil
}

func RecordContainerInfo(pid, name, id, command string, volumes []string, rootPath string) error {
	containerInfo := &ContainerInfo {
		Pid: 		pid,
		Id:  		id,
		Name:		name,
		Command: 	command,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		Status:		RUNNING,
		Volumes: 	volumes,
		RootPath: 	rootPath,
	}
	jsonInfo, _ := json.Marshal(containerInfo)
	//log.Printf("jsonInfo:%s\n", string(jsonInfo))
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
	//str := time.Now().UnixNano()
	//containerId := fmt.Sprintf("%d%d", str, int(math.Abs(float64(rand.Intn(10)))))
	////log.Printf("containerId:%s\n", containerId)
	//return containerId
	letterBytes := "1234567890"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetLogFile(containerName string) (*os.File, error) {
	path := fmt.Sprintf(INFOLOCATION, containerName)
	logFile := path + "/" + CONTAINERLOGS
	if err := os.MkdirAll(path, 0622); err != nil {
		return nil, fmt.Errorf("create %s error : %v\n", path, err)
	}
	if file , err := os.Create(logFile); err != nil {
		return nil, fmt.Errorf("os.Create(%s) error : %v\n", path, err)
	} else {
		return file, nil
	}
}

func ReadLogs(containerName string) string {
	path := fmt.Sprintf(INFOLOCATION, containerName)
	logFile := path + "/" + CONTAINERLOGS
	data, _ := ioutil.ReadFile(logFile)
	return string(data)
}

func UpdateContainerInfo(containerInfo *ContainerInfo) error {
	jsonInfo, _ := json.Marshal(containerInfo)
	//log.Printf("jsonInfo:%s\n", string(jsonInfo))
	location := fmt.Sprintf(INFOLOCATION, containerInfo.Name)
	file 	 := location + "/" + CONFIGNAME
	if err := ioutil.WriteFile(file, []byte(jsonInfo), 0622); err != nil {
		return fmt.Errorf("write %s to %s error:%v\n", jsonInfo, file, err)
	}
	return nil
}

func RemoveContainerInfo(containerInfo *ContainerInfo) error {
	location := fmt.Sprintf(INFOLOCATION, containerInfo.Name)
	if err := os.RemoveAll(location); err != nil {
		return fmt.Errorf("os.RemoveAll(%s) %v\n", location, err)
	}
	return nil
}