package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/x87w/zenith/pkg/storage"
	"github.com/x87w/zenith/pkg/vm"
)

const vmDir = "vms"

func Run() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {
	case "create":
		create()
	case "start":
		start()
	case "list":
		list()
	default:
		usage()
	}
}

func create() {
	fs := flag.NewFlagSet("create", flag.ExitOnError)

	name := fs.String("name", "", "VM name")
	memory := fs.String("memory", "2G", "Memory")
	cpu := fs.Int("cpu", 2, "CPU cores")
	size := fs.String("size", "20G", "Disk size")
	iso := fs.String("iso", "", "Installation ISO")

	fs.Parse(os.Args[2:])

	if *name == "" {
		fmt.Println("missing required flag: --name")
		return
	}

	if err := os.MkdirAll(vmDir, 0755); err != nil {
		fmt.Println(err)
		return
	}

	path := filepath.Join(vmDir, *name+".qcow2")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := storage.CreateQCOW2(path, *size); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("created", path)
	}

	cfg := vm.Config{
		Name:   *name,
		Memory: *memory,
		CPUs:   *cpu,
		Disk:   path,
		ISO:    *iso,
	}

	fmt.Printf("Starting %s...\n", cfg.Name)

	if err := vm.New(cfg).Start(); err != nil {
		fmt.Println(err)
	}
}

func start() {
	fs := flag.NewFlagSet("start", flag.ExitOnError)

	name := fs.String("name", "", "VM name")
	memory := fs.String("memory", "2G", "Memory")
	cpu := fs.Int("cpu", 2, "CPU cores")

	fs.Parse(os.Args[2:])

	if *name == "" {
		fmt.Println("missing required flag: --name")
		return
	}

	path := filepath.Join(vmDir, *name+".qcow2")

	if _, err := os.Stat(path); err != nil {
		fmt.Println("VM not found:", *name)
		return
	}

	cfg := vm.Config{
		Name:   *name,
		Memory: *memory,
		CPUs:   *cpu,
		Disk:   path,
	}

	fmt.Printf("Starting %s...\n", cfg.Name)

	if err := vm.New(cfg).Start(); err != nil {
		fmt.Println(err)
	}
}

func list() {
	if err := os.MkdirAll(vmDir, 0755); err != nil {
		fmt.Println(err)
		return
	}

	entries, err := os.ReadDir(vmDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(entries) == 0 {
		fmt.Println("No VMs found.")
		return
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
}

func usage() {
	fmt.Println("Zenith VM Manager")
	fmt.Println()

	fmt.Println("Commands:")
	fmt.Println("  create --name <name> [options]")
	fmt.Println("  start  --name <name>")
	fmt.Println("  list")
	fmt.Println()

	fmt.Println("Options:")
	fmt.Println("  --memory 2G")
	fmt.Println("  --cpu 2")
	fmt.Println("  --size 20G")
	fmt.Println("  --iso <path>")
}
