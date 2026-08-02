package vm

import (
	"fmt"
	"os"
	"os/exec"
)

// Config holds startup params for the vm
type Config struct {
	Name   string
	Memory string
	CPUs   int
	Disk   string
}

type VM struct {
	Config Config
	cmd    *exec.Cmd
}

// New constructs a new vm reference
func New(cfg Config) *VM {
	return &VM{Config: cfg}
}

// Start executes qemu-system-x86_64 with the given configs
func (v *VM) Start() error {
	args := []string{
		"-name", v.Config.Name,
		"-m", v.Config.Memory,
		"-smp", fmt.Sprintf("%d", v.Config.CPUs),
		"-drive", fmt.Sprintf("file=%s,format=qcow2", v.Config.Disk),
		"-nographic",
	}

	v.cmd = exec.Command("qemu-system-x86_64", args...)
	v.cmd.Stdout = os.Stdout
	v.cmd.Stderr = os.Stderr
	v.cmd.Stdin = os.Stdin

	return v.cmd.Start()
}
