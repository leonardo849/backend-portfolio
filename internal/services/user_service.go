package services

import (
	"backend-portfolio/internal/dto"
	"backend-portfolio/internal/helper"
	"backend-portfolio/internal/logger"
	"backend-portfolio/internal/repository"
	"backend-portfolio/internal/validate"
	"go.uber.org/zap"
)

type UserService struct {
	userRepository *repository.UserRepository
	prefix string
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return  &UserService{
		userRepository: userRepository,
		prefix: "user ",
	}
}


func (u *UserService) LoginUser(input dto.LoginUserDTO) (status int, message interface{}) {
	if err := validate.Validate.Struct(input); err != nil {
		logger.ZapLogger.Error(
			"validate error",
			zap.Error(err),
			zap.String("function", "login user service"),
		)
		return 400, err.Error()
	}
	res := u.userRepository.LoginUser(input)
	
	if  res.Error != nil {
		code := helper.HandleErrors(res.Error)
		mes := u.prefix + res.Error.Error()
		return *code, mes
	} 

	return 200, res.Data
	
}