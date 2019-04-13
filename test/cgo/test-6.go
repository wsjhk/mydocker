package main

import (
	"log"
	_ "github.com/nicktming/mydocker/test/ccode"
	"os"
)

func main()  {
	mydocker_pid := os.Args[0]
	log.Printf("from args mydocker_pid:%s\n", mydocker_pid)
	os.Setenv("mydocker_pid", mydocker_pid)
	log.Printf("hello world!")
}