package cgroups

import "github.com/nicktming/mydocker/cgroups/subsystems"

type CroupManger struct {
	// 限制的值
	Resource *subsystems.ResourceConfig
	// 用于在当前容器中标识有哪些subsystem需要做限制
	SubsystemsIns []subsystems.Subsystem
	// 用于启动多个container的containerId 比如memory则是:/sys/fs/cgroup/memory/mydocker/[containerId]
	//path string
}

func (c *CroupManger) Set() {
	for _, sub := range c.SubsystemsIns {
		sub.Set(c.Resource)
	}
}

func (c *CroupManger) Apply(pid string) {
	for _, sub := range c.SubsystemsIns {
		sub.Apply(pid)
	}
}

func (c *CroupManger) Destroy() {
	for _, sub := range c.SubsystemsIns {
		sub.Remove()
	}
}
