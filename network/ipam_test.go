package network

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"path"
	"strings"
	"testing"
)

func Test001(t *testing.T)  {
	err := ipAllocator.dump()
	log.Printf("error:%v\n", err)
}

func Test002(t *testing.T)  {
	hostip, ipnet, _ := net.ParseCIDR("192.168.0.1/24")
	log.Printf("ip: %s ipnet ip:%s, mask:%s\n", hostip, ipnet.IP, ipnet.Mask)
	ip, _ := ipAllocator.Allocate(ipnet)
	log.Printf("alloc ip : %v\n", ip)
}

func Test003(t *testing.T)  {
	hostip, ipnet, _ := net.ParseCIDR("192.168.0.1/24")
	log.Printf("ip: %s ipnet ip:%s, mask:%s\n", hostip, ipnet.IP, ipnet.Mask)
	ipAllocator.Release(ipnet, &hostip)
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

	log.Printf("n:%d\n", n)
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
	log.Printf("dump ipamConfigFileDir:%s\n", ipamConfigFileDir)
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

	log.Printf("dump ipamConfigJson:%s\n", ipamConfigJson)

	_, err = subnetConfigFile.Write(ipamConfigJson)
	if err != nil {
		return err
	}

	return nil
}

func (ipam *IPAM) Allocate(subnet *net.IPNet) (ip net.IP, err error) {
	// 存放网段中地址分配信息的数组
	ipam.Subnets = &map[string]string{}

	// 从文件中加载已经分配的网段信息
	err = ipam.load()
	if err != nil {
		log.Printf("Error dump allocation info, %v", err)
	}

	_, subnet, _ = net.ParseCIDR(subnet.String())

	log.Printf("Allocate subnet:%s, ipam.Subnets:%v\n", subnet, ipam.Subnets)

	one, size := subnet.Mask.Size()

	log.Printf("Allocate one:%d, size:%d\n", one, size)

	if _, exist := (*ipam.Subnets)[subnet.String()]; !exist {
		(*ipam.Subnets)[subnet.String()] = strings.Repeat("0", 1 << uint8(size - one))
	}

	log.Printf("Allocate one:%s\n", (*ipam.Subnets)[subnet.String()])

	for c := range((*ipam.Subnets)[subnet.String()]) {
		if (*ipam.Subnets)[subnet.String()][c] == '0' {
			ipalloc := []byte((*ipam.Subnets)[subnet.String()])
			ipalloc[c] = '1'
			(*ipam.Subnets)[subnet.String()] = string(ipalloc)
			ip = subnet.IP
			for t := uint(4); t > 0; t-=1 {
				[]byte(ip)[4-t] += uint8(c >> ((t - 1) * 8))
			}
			ip[3]+=1
			break
		}
	}

	ipam.dump()
	return
}

func (ipam *IPAM) Release(subnet *net.IPNet, ipaddr *net.IP) error {
	ipam.Subnets = &map[string]string{}

	_, subnet, _ = net.ParseCIDR(subnet.String())

	err := ipam.load()
	if err != nil {
		log.Printf("Error dump allocation info, %v", err)
	}

	c := 0
	releaseIP := ipaddr.To4()
	releaseIP[3]-=1
	for t := uint(4); t > 0; t-=1 {
		c += int(releaseIP[t-1] - subnet.IP[t-1]) << ((4-t) * 8)
	}

	ipalloc := []byte((*ipam.Subnets)[subnet.String()])
	ipalloc[c] = '0'
	(*ipam.Subnets)[subnet.String()] = string(ipalloc)

	ipam.dump()
	return nil
}