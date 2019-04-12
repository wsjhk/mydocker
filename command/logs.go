package command

import "fmt"

func Logs(containerName string) {
	data := ReadLogs(containerName)
	fmt.Println(data)
}
