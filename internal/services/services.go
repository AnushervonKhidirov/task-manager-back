package services

import (
	"encoding/json"

	"main/internal/database"
	"main/internal/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	var user model.User
	unmarshalErr := json.Unmarshal(c.Body(), &user)

	if unmarshalErr != nil {
		return unmarshalErr
	}

	hashedPassword, hashingErr := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if hashingErr != nil {
		return hashingErr
	}

	user.Password = string(hashedPassword)

	result := database.AddUser(&user)

	responseBody, marshalErr := json.Marshal(result)

	if marshalErr != nil {
		return marshalErr
	}

	c.Send(responseBody)

	return nil
}
