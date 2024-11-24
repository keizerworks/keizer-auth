package controllers

import (
	"keizer-auth/internal/models"
	"keizer-auth/internal/services"
	"keizer-auth/internal/utils"
	"keizer-auth/internal/validators"

	"github.com/gofiber/fiber/v2"
)

type ApplicationController struct {
	applicationService *services.ApplicationService
}

func NewApplicationController(
	applicationService *services.ApplicationService,
) *ApplicationController {
	return &ApplicationController{
		applicationService: applicationService,
	}
}

func (self *ApplicationController) Get(c *fiber.Ctx) error {
	user := utils.GetCurrentUser(c)
	applications, err := self.applicationService.Get(user.ID, user.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(applications)
}

func (self *ApplicationController) Create(c *fiber.Ctx) error {
	var err error
	user := utils.GetCurrentUser(c)
	body := new(validators.CreateApplication)

	if err := c.BodyParser(body); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := body.ValidateFile(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	account := new(models.Account)
	account, err = self.applicationService.Create(body.Name, account.ID user.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(account)
}
