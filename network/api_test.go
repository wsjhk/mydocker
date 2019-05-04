package network

import (
	"fmt"
	"log"
	"net"
	"strings"
	"testing"
	"github.com/vishvananda/netlink"
	"time"
)


// 类似于 ip link add name testbridge type bridge
func TestNet001(t *testing.T) {
	bridgeName := "testbridge"
	_, err := net.InterfaceByName(bridgeName)
	if err == nil || !strings.Contains(err.Error(), "no such network interface") {
		log.Printf("error:%v\n", err)
	}
	// create *netlink.Bridge object
	la := netlink.NewLinkAttrs()
	la.Name = bridgeName

	br := &netlink.Bridge{la}
	if err := netlink.LinkAdd(br); err != nil {
		fmt.Errorf("Bridge creation failed for bridge %s: %v", bridgeName, err)
	}
}

func TestNet002(t *testing.T) {
	name := "testbridge"
	rawIP := "192.168.0.1/24"
	retries := 2
	var iface netlink.Link
	var err error
	for i := 0; i < retries; i++ {
		iface, err = netlink.LinkByName(name)
		if err == nil {
			break
		}
		log.Printf("error retrieving new bridge netlink link [ %s ]... retrying", name)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		fmt.Errorf("Abandoning retrieving the new bridge link from netlink, Run [ ip link ] to troubleshoot the error: %v", err)
	}
	ipNet, err := netlink.ParseIPNet(rawIP)
	if err != nil {
		log.Printf("ParseIPNet error:%v\n", err)
	}
	log.Printf("ipNet:%v\n", ipNet)
	addr := &netlink.Addr{ipNet, "", 0, 0, nil}
	err = netlink.AddrAdd(iface, addr)
	log.Printf("AddrAdd error:%v\n", err)

	if err := netlink.LinkSetUp(iface); err != nil {
		fmt.Errorf("Error enabling interface for %s: %v", name, err)
	}
}
