package user

import "github.com/gofiber/fiber/v2"

func User() *fiber.App {
	user := fiber.New()
	{
		user.Post("/addUser", AddUser)
	}
	return user
}

func AddUser(c *fiber.Ctx) (err error) {

	return
}
