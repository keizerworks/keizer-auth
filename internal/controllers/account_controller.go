package controllers

import (
	"keizer-auth/internal/models"
	"keizer-auth/internal/services"
	"keizer-auth/internal/utils"
	"keizer-auth/internal/validators"

	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	accountService *services.AccountService
}

func NewAccountController(
	accountService *services.AccountService,
) *AccountController {
	return &AccountController{accountService: accountService}
}

func (self *AccountController) Get(c *fiber.Ctx) error {
	user := utils.GetCurrentUser(c)
	accounts, err := self.accountService.GetAccountsByUser(user.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(accounts)
}

func (self *AccountController) Create(c *fiber.Ctx) error {
	var err error
	user := utils.GetCurrentUser(c)
	body := new(validators.CreateAccount)

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
	account, err = self.accountService.Create(body.Name, user.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(account)
}
