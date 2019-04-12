package command

import "fmt"

func Logs(containerName string) error {
	data := ReadLogs(containerName)
	fmt.Println(data)
}
