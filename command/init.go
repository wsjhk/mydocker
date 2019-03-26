package command

import (
	"log"
	"os"
	"syscall"
)

func Init(command string)  {

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")

	/*
	cmd := exec.Command(command)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Printf("Init Run() function err : %v\n", err)
		log.Fatal(err)
	}
	*/

	if err := syscall.Exec(command, []string{command}, os.Environ()); err != nil {
		log.Printf("syscall.Exec err: %v\n", err)
		log.Fatal(err)
	}
}
