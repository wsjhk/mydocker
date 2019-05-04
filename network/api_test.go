package network

import (
	"fmt"
	"log"
	"net"
	"strings"
	"testing"
	"github.com/vishvananda/netlink"
)

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
