package handlers

import (
	"backend-portfolio/internal/dto"
	"backend-portfolio/internal/helper"
	"backend-portfolio/internal/logger"
	"backend-portfolio/internal/services"
	"go.uber.org/zap"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController{
	return  &UserController{
		userService: userService,
	}
}

// LoginUser godoc
// @Description Login user
// @Tags users
// @Accept json
// @Produce json
// @Param guest body dto.LoginUserDTO true "user data"
// @Success 201 {object} dto.MessageResponseDTO
// @Failure 401 {object} dto.ErrorResponseDTO
// @Router /user/login [post]
func (u *UserController) LoginUser() fiber.Handler {
	return func (ctx *fiber.Ctx) error {
		var input dto.LoginUserDTO
		if err := ctx.BodyParser(&input); err != nil {
			logger.ZapLogger.Error("request body to dto.loginuserdto", zap.Error(err), zap.String("function", "login user controller"))
			return  ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		status, message := u.userService.LoginUser(input)
		property := helper.SetProperty(status)
		return  ctx.Status(status).JSON(fiber.Map{property: message})
	}
}