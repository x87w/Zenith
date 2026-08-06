package storage

import (
	"os"
	"os/exec"
)

func CreateQCOW2(path string, size string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	}

	return exec.Command(
		"qemu-img",
		"create",
		"-f",
		"qcow2",
		path,
		size,
	).Run()
}
