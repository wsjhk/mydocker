package network

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"testing"
)

func Test001(t *testing.T)  {
	err := ipAllocator.dump()
	log.Printf("error:%v\n", err)
}

const ipamDefaultAllocatorPath = "/var/run/mydocker/network/ipam/subnet.json"

type IPAM struct {
	SubnetAllocatorPath string
	Subnets *map[string]string
}

var ipAllocator = &IPAM{
	SubnetAllocatorPath: ipamDefaultAllocatorPath,
}

func (ipam *IPAM) load() error {
	if _, err := os.Stat(ipam.SubnetAllocatorPath); err != nil {
		log.Printf("load error err:%v\n", err)
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	}
	subnetConfigFile, err := os.Open(ipam.SubnetAllocatorPath)
	defer subnetConfigFile.Close()
	if err != nil {
		return err
	}
	subnetJson := make([]byte, 2000)
	n, err := subnetConfigFile.Read(subnetJson)
	if err != nil {
		return err
	}

	log.Println(subnetJson)

	err = json.Unmarshal(subnetJson[:n], ipam.Subnets)
	if err != nil {
		log.Printf("Error dump allocation info, %v", err)
		return err
	}
	return nil
}

func (ipam *IPAM) dump() error {
	ipamConfigFileDir, _ := path.Split(ipam.SubnetAllocatorPath)
	log.Printf("ipamConfigFileDir:%s\n", ipamConfigFileDir)
	if _, err := os.Stat(ipamConfigFileDir); err != nil {
		if os.IsNotExist(err) {
			log.Printf("MkdirAll\n")
			os.MkdirAll(ipamConfigFileDir, 0644)
		} else {
			return err
		}
	}
	// O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	// O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	// O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	subnetConfigFile, err := os.OpenFile(ipam.SubnetAllocatorPath, os.O_TRUNC | os.O_WRONLY | os.O_CREATE, 0644)
	defer subnetConfigFile.Close()
	if err != nil {
		return err
	}

	ipamConfigJson, err := json.Marshal(ipam.Subnets)
	if err != nil {
		return err
	}

	_, err = subnetConfigFile.Write(ipamConfigJson)
	if err != nil {
		return err
	}

	return nil
}
