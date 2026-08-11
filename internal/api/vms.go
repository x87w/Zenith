package api

import (
	"path/filepath"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"

	"github.com/x87w/zenith/pkg/storage"
	"github.com/x87w/zenith/pkg/vm"
)

func CreateVM(c fiber.Ctx) error {
	type request struct {
		Name   string `json:"name"`
		Memory string `json:"memory"`
		CPU    int    `json:"cpu"`
		Disk   string `json:"disk"`
	}

	var body request

	if err := c.Bind().Body(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	if body.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "name required",
		})
	}

	if body.Memory == "" {
		body.Memory = "2G"
	}

	if body.CPU == 0 {
		body.CPU = 2
	}

	id := uuid.New().String()

	path := filepath.Join("vms", id+".qcow2")

	if err := storage.CreateQCOW2(path, "20G"); err != nil {
		return err
	}

	cfg := vm.Config{
		Name:   body.Name,
		Memory: body.Memory,
		CPUs:   body.CPU,
		Disk:   path,
	}

	go vm.New(cfg).Start()

	return c.JSON(fiber.Map{
		"id":     id,
		"name":   body.Name,
		"status": "starting",
	})
}
