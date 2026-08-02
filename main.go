package main

import (
	"fmt"
	"log"

	"github.com/x87w/zenith/pkg/vm"
)

func main() {
	fmt.Println("zenith virt manager")

	cfg := vm.Config{
		Name:   "zenith-node-01",
		Memory: "2048M",
		CPUs:   2,
		Disk:   "./disk.qcow2",
	}

	machine := vm.New(cfg)

	fmt.Printf("Booting %s...\n", cfg.Name)
	if err := machine.Start(); err != nil {
		log.Fatalf("Error launching VM: %v\n", err)
	}

	fmt.Println("VM process launched successfully!")
}
