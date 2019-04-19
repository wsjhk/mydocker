package main

import (
	"github.com/nicktming/mydocker/command"
	"log"
	"os/exec"
)

func main()  {
	test01()
}

func test01()  {
	copy("/root/filecopy/busybox", "/root/filecopy/tmp")
	copy("/root/filecopy/busybox/bin/top", "/root/filecopy/tmp")
}

func copy(src, dst string) {
	exist, _ := command.PathExists(src)
	if !exist {
		log.Printf("src:%s not exists!\n", src)
		return
	}
	exist, _ = command.PathExists(dst)
	if !exist {
		log.Printf("dst:%s not exists!\n", src)
		return
	}
	if _, err := exec.Command("cp", "-r", src, dst).CombinedOutput(); err != nil {
		log.Printf("cp -r %s %s, err:%v\n", src, dst, err)
		return
	}
}