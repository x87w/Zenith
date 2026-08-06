package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/x87w/zenith/internal/db"
)

func Register(c fiber.Ctx) error {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var body request

	if err := c.Bind().Body(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(body.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := db.User{
		ID:       uuid.New().String(),
		Username: body.Username,
		Password: string(hash),
		Token:    uuid.New().String(),
	}

	db.Users = append(db.Users, user)

	return c.JSON(fiber.Map{
		"id":    user.ID,
		"token": user.Token,
	})
}
