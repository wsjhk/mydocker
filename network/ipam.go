package network

//import (
//	"encoding/json"
//	"log"
//	"os"
//)
//
//const ipamDefaultAllocatorPath = "/var/run/mydocker/network/ipam/subnet.json"
//
//type IPAM struct {
//	SubnetAllocatorPath string
//	Subnets *map[string]string
//}
//
//var ipAllocator = &IPAM{
//	SubnetAllocatorPath: ipamDefaultAllocatorPath,
//}
//
//func (ipam *IPAM) load() error {
//	if _, err := os.Stat(ipam.SubnetAllocatorPath); err != nil {
//		if os.IsNotExist(err) {
//			return nil
//		} else {
//			return err
//		}
//	}
//	subnetConfigFile, err := os.Open(ipam.SubnetAllocatorPath)
//	defer subnetConfigFile.Close()
//	if err != nil {
//		return err
//	}
//	subnetJson := make([]byte, 2000)
//	n, err := subnetConfigFile.Read(subnetJson)
//	if err != nil {
//		return err
//	}
//
//	log.Println(subnetJson)
//
//	err = json.Unmarshal(subnetJson[:n], ipam.Subnets)
//	if err != nil {
//		log.Printf("Error dump allocation info, %v", err)
//		return err
//	}
//	return nil
//}