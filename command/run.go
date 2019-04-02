package command

import (
	"github.com/nicktming/mydocker/cgroups"
	"log"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func Run(command string, tty bool, cg *cgroups.CroupManger)  {
	//cmd := exec.Command(command)

	cmd := exec.Command("/proc/self/exe", "init", command)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if tty {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
	}
	/**
	 *   Start() will not block, so it needs to use Wait()
	 *   Run() will block
	 */
	if err := cmd.Start(); err != nil {
		log.Printf("Run Start err: %v.\n", err)
		log.Fatal(err)
	}
	//log.Printf("222 before process pid:%d, memory:%s\n", cmd.Process.Pid, memory)

	//subsystems.Set(memory)
	//subsystems.Apply(strconv.Itoa(cmd.Process.Pid))
	//defer subsystems.Remove()

	cg.Set()
	defer cg.Destroy()
	cg.Apply(strconv.Itoa(cmd.Process.Pid))

	cmd.Wait()
}
