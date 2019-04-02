package cgroups

import "github.com/nicktming/mydocker/cgroups/subsystems"



type CgroupManger struct {
	// 限制的值
	Resource *subsystems.ResourceConfig
	// 用于在当前容器中标识有哪些subsystem需要做限制
	SubsystemsIns []subsystems.Subsystem
	// 用于启动多个container的containerId 比如memory则是:/sys/fs/cgroup/memory/mydocker/[containerId]
	//path string
}

func (c *CgroupManger) Set() {
	for _, sub := range c.SubsystemsIns {
		sub.Set(c.Resource)
	}
}

func (c *CgroupManger) Apply(pid string) {
	for _, sub := range c.SubsystemsIns {
		sub.Apply(pid)
	}
}

func (c *CgroupManger) Destroy() {
	for _, sub := range c.SubsystemsIns {
		sub.Remove()
	}
}
