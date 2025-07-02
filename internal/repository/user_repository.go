package repository

import (
	"backend-portfolio/internal/dto"
	"backend-portfolio/internal/helper"
	"backend-portfolio/internal/logger"
	"backend-portfolio/internal/models"
	"errors"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type repositoryResponse[T any] struct {
	Data T
	Error error
}



type UserRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(userCollection *mongo.Collection) *UserRepository {
	return  &UserRepository{
		userCollection: userCollection,
	}
}

func (u *UserRepository) FindOneUserByUsername(username string, omitPassword bool) repositoryResponse[*dto.FindUserDTO] {
	ctx, cancel := helper.CreateCtx()
	defer cancel()
	var user models.UserModel
	err := u.userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.ZapLogger.Error("user not found", zap.Error(err), zap.String("function", "find one user by username repo"))
			return  repositoryResponse[*dto.FindUserDTO]{
				Data: nil,
				Error: fmt.Errorf(helper.NOTFOUND),
			}
		} else {
			logger.ZapLogger.Error("internal server error in find one", zap.Error(err), zap.String("function", "find one user by username repo"))
			return repositoryResponse[*dto.FindUserDTO]{
				Data: nil,
				Error: fmt.Errorf(helper.INTERNALSERVER),
			}
		}
	}
	userDTO := dto.FindUserDTO{
		ID: user.ID,
		PhotoURL: user.PhotoURL,
		Username: user.Username,
		Slogan: user.Slogan,
		Description: user.Description,
		Skills: user.Skills,
		Formations: nil,
		SocialMedias: nil,
		Role: user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Password: nil,
	}
	if (!omitPassword) {
		userDTO.Password = &user.Password
	} 
	message := "find one was made " + "ommited password:" + strconv.FormatBool(omitPassword) + " searched id:" + user.ID.String() 
	logger.ZapLogger.Info(message)
	return  repositoryResponse[*dto.FindUserDTO]{
			Data: &userDTO,
			Error: nil,
	}
}

func (u *UserRepository) LoginUser(input dto.LoginUserDTO) repositoryResponse[*string] {
	response := u.FindOneUserByUsername(input.Username, false)
	if response.Error != nil && response.Error.Error() == helper.NOTFOUND {
		return repositoryResponse[*string]{
			Data: nil,
			Error: response.Error,
		}
	}
	userDto := response.Data
	arePasswordsEqual := helper.CompareHash(input.Password, *userDto.Password)
	if arePasswordsEqual {
		jwt, err := helper.CreateJWT(userDto.ID.String(), userDto.Username, userDto.Role ,userDto.UpdatedAt)
		if err != nil {
			logger.ZapLogger.Error("error in create jwt", 
			zap.Error(err),
			zap.String("function", "login user repository"),
			)
			return repositoryResponse[*string]{
				Error: fmt.Errorf(helper.INTERNALSERVER),
				Data: nil,
			} 
		}
		message := "user with id: " + userDto.ID.String() + " " + "logged in the system"
		logger.ZapLogger.Info(message)
		return repositoryResponse[*string]{
			Error: nil,
			Data: &jwt,
		}
	} 
	logger.ZapLogger.Error("login failed. Credentials are wrong", zap.String("function", "login user repository"))
	return repositoryResponse[*string]{
		Data: nil,
		Error: fmt.Errorf("your credentials are wrong"),
	}

}