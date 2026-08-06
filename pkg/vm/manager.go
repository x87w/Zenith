package vm

import (
	"fmt"
	"os"
	"os/exec"
)

type Config struct {
	Name   string
	Memory string
	CPUs   int
	Disk   string
	ISO    string
}

type VM struct {
	Config Config
	cmd    *exec.Cmd
}

func New(cfg Config) *VM {
	return &VM{Config: cfg}
}

func (v *VM) Start() error {
	args := []string{
		"-machine", "q35,accel=hvf",
		"-cpu", "host",
		"-name", v.Config.Name,
		"-m", v.Config.Memory,
		"-smp", fmt.Sprintf("%d", v.Config.CPUs),

		"-drive", fmt.Sprintf("file=%s,if=virtio,format=qcow2", v.Config.Disk),

		"-netdev", "user,id=net0,hostfwd=tcp::2222-:22",
		"-device", "virtio-net-pci,netdev=net0",

		"-serial", "mon:stdio",
		"-nographic",
	}

	if v.Config.ISO != "" {
		args = append(args,
			"-cdrom", v.Config.ISO,
			"-boot", "d",
		)
	}

	fmt.Println("qemu args:", args)

	v.cmd = exec.Command("qemu-system-x86_64", args...)
	v.cmd.Stdout = os.Stdout
	v.cmd.Stderr = os.Stderr
	v.cmd.Stdin = os.Stdin

	return v.cmd.Run()
}
