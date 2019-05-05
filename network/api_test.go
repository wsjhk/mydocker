package network

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netns"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"
)

// 等于 ip link add name testbridge type bridge
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
	// 等于 ip link add name testbridge type bridge
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
		// 根据名字找到设备
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
	// 将原始ip转换成*net.IPNet类型
	ipNet, err := netlink.ParseIPNet(rawIP)
	if err != nil {
		log.Printf("ParseIPNet error:%v\n", err)
	}

	log.Printf("ipNet:%v\n", ipNet)
	addr := &netlink.Addr{ipNet, "", 0, 0, nil}

	// 等于 ip addr add 192.168.0.1/24 dev testbridge
	err = netlink.AddrAdd(iface, addr)
	log.Printf("AddrAdd error:%v\n", err)

	// 等于 ip link set testbridge up
	if err := netlink.LinkSetUp(iface); err != nil {
		fmt.Errorf("Error enabling interface for %s: %v", name, err)
	}
}

func TestNet003(t *testing.T) {
	bridgeName := "testbridge"
	// 根据设备名找到设备testbridge
	br, err := netlink.LinkByName(bridgeName)
	if err != nil {
		log.Printf("LinkByName err:%v\n", err)
		return
	}

	la := netlink.NewLinkAttrs()
	la.Name = "12345"

	log.Printf("br.attrs().index:%d\n", br.Attrs().Index)
	// 等于 ip link set dev 12345 master testbridge
	la.MasterIndex = br.Attrs().Index

	myVeth := netlink.Veth{
		LinkAttrs: la,
		PeerName:  "cif-" + la.Name,
	}
	// 等于 ip link add 12345 type veth peer name cif-12345
	if err = netlink.LinkAdd(&myVeth); err != nil {
		fmt.Errorf("Error Add Endpoint Device: %v", err)
		return
	}

	// 等于 ip link set 12345 up
	if err = netlink.LinkSetUp(&myVeth); err != nil {
		fmt.Errorf("Error Add Endpoint Device: %v", err)
		return
	}
}

func TestNet005(t *testing.T) {
	deleteDevice("testbridge")
	deleteDevice("12345")
}

func deleteDevice(name string)  {
	// 根据设备名找到该设备
	l, err := netlink.LinkByName(name)
	if err != nil {
		fmt.Errorf("Getting link with name %s failed: %v", name, err)
		return
	}

	// 删除设备
	// 删除网桥就等于 ifconfig testbridge down && ip link delete testbridge type bridge
	// 删除veth就等于  ip link delete 12345 type veth
	if err := netlink.LinkDel(l); err != nil {
		fmt.Errorf("Failed to remove bridge interface %s delete: %v", name, err)
		return
	}
	log.Printf("Delete Device %s\n", name)
}

func TestNet006(t *testing.T) {
	PeerName := "cif-12345"
	containerIP := "192.168.0.8/24"

	gwIP, ipnet, _ := net.ParseCIDR("192.168.0.1/24")
	ipnet.IP = gwIP


	if err := ConfigEndpointIpAddressAndRoute(PeerName, containerIP, gwIP); err != nil {
		log.Printf("ConfigEndpointIpAddressAndRoute error:%v\n", err)
	}
}


func EnterContainerNetns(enLink *netlink.Link) func() {
	f, err := os.OpenFile(fmt.Sprintf("/proc/%s/ns/net", "16694"), os.O_RDONLY, 0)
	if err != nil {
		logrus.Errorf("error get container net namespace, %v", err)
	}

	nsFD := f.Fd()
	runtime.LockOSThread()

	// 修改veth peer 另外一端移到容器的namespace中
	if err = netlink.LinkSetNsFd(*enLink, int(nsFD)); err != nil {
		logrus.Errorf("error set link netns , %v", err)
	}

	// 获取当前的网络namespace
	origns, err := netns.Get()
	if err != nil {
		logrus.Errorf("error get current netns, %v", err)
	}

	// 设置当前进程到新的网络namespace，并在函数执行完成之后再恢复到之前的namespace
	if err = netns.Set(netns.NsHandle(nsFD)); err != nil {
		logrus.Errorf("error set netns, %v", err)
	}
	return func () {
		netns.Set(origns)
		origns.Close()
		runtime.UnlockOSThread()
		f.Close()
	}
}

func ConfigEndpointIpAddressAndRoute(PeerName, containerIP string, ipnet net.IPNet) error {
	peerLink, err := netlink.LinkByName(PeerName)
	if err != nil {
		return fmt.Errorf("fail config endpoint: %v", err)
	}

	defer EnterContainerNetns(&peerLink)()

	if err = setInterfaceIP(PeerName, containerIP); err != nil {
		return fmt.Errorf("set %s up error:%v", PeerName, err)
	}

	if err = setInterfaceUP(PeerName); err != nil {
		return err
	}

	if err = setInterfaceUP("lo"); err != nil {
		return err
	}

	_, cidr, _ := net.ParseCIDR("0.0.0.0/0")

	defaultRoute := &netlink.Route{
		LinkIndex: peerLink.Attrs().Index,
		Gw: ipnet.IP,
		Dst: cidr,
	}

	if err = netlink.RouteAdd(defaultRoute); err != nil {
		return err
	}

	return nil
}

func setInterfaceUP(interfaceName string) error {
	iface, err := netlink.LinkByName(interfaceName)
	if err != nil {
		return fmt.Errorf("Error retrieving a link named [ %s ]: %v", iface.Attrs().Name, err)
	}

	if err := netlink.LinkSetUp(iface); err != nil {
		return fmt.Errorf("Error enabling interface for %s: %v", interfaceName, err)
	}
	return nil
}

// Set the IP addr of a netlink interface
func setInterfaceIP(name string, rawIP string) error {
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
		return fmt.Errorf("Abandoning retrieving the new bridge link from netlink, Run [ ip link ] to troubleshoot the error: %v", err)
	}
	ipNet, err := netlink.ParseIPNet(rawIP)
	if err != nil {
		return err
	}
	addr := &netlink.Addr{ipNet, "", 0, 0, nil}
	return netlink.AddrAdd(iface, addr)
}
