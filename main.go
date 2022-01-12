package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/sonerrtng/go-social-media/entity"
	"github.com/sonerrtng/go-social-media/repositories"
)

func main() {

	app := fiber.New()

	// Routes
	app.Post("/v1/user", createUser)

	app.Get("/v1/user", getUser)

	// Start server
	app.Listen(":3000")
}

func createUser(c *fiber.Ctx) error {

	userInfo := entity.User{}
	json.Unmarshal(c.Body(), &userInfo)
	control := userInfo.ControlUserInfo()

	if !control {
		return c.SendString("{'status':'error','error':'missing or incorrect information'}")
	}

	mysqlRepo := repositories.CreateMysqlRepository()
	userInsertStatus := mysqlRepo.CreateUser(userInfo)
	if userInsertStatus == -1 {
		return c.SendString("{'status':'error','error':'missing or incorrect information'}")
	}
	return c.SendString("{'status':'succsesfully'}")
}

func getUser(c *fiber.Ctx) error {
	mysqlRepo := repositories.CreateMysqlRepository()
	users := mysqlRepo.GetUser()

	usersJson, _ := json.Marshal(users)

	return c.SendString(string(usersJson))
}
