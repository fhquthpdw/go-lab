package main

// NOT FINISHED
import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker run <container> cmd args
// go run main.go run cmd args
func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("what??")
	}
}

func run() {
	fmt.Printf("running %v\n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Env = []string{"a=b", "a=c"}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | // UTS(UNIX Time Sharing，主机名与域名)
			syscall.CLONE_NEWIPC | // 信号量，消息对列和共享内存
			syscall.CLONE_NEWNS | // 挂载点，文件系统
			syscall.CLONE_NEWPID | // PID 进程编号
			syscall.CLONE_NEWNET | // 网络设备，网络栈和端口
			syscall.CLONE_NEWUSER, // 用户和用户组
	}

	must(syscall.Sethostname([]byte("inside-container"))) // 设置主机名
	must(syscall.Chroot("/home/rootfs"))                  // 设置根目录，文件系统的根目录: /
	must(os.Chdir("/"))                                   // 设置工作目录

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
