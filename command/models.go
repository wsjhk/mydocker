package command


type ContainerInfo struct {
	Pid			string	`json:"pid"`
	Id			string	`json:"id"`
	Name		string	`json:"name"`
	Command		string	`json:"command"`
	CreateTime	string	`json:"createTime"`
	Status		string	`json:"status"`
}

var (
	RUNNING			 = "running"
	STOP			 = "stopped"
	EXIT			 = "exited"
	CONTAINS         = "/var/run/mydocker"
	INFOLOCATION	 = "/var/run/mydocker/%s"
	CONFIGNAME		 = "config.json"
	CONTAINERLOGS	 = "container.log"
)
