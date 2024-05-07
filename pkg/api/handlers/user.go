package handlers

import (
	"net/http"
	services "scalexmedia-assignment/pkg/usecase/interface"
	"scalexmedia-assignment/pkg/utils/models"
	"scalexmedia-assignment/pkg/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	UserUseCase services.UserUseCase
	userMap     map[string]models.UserStruct
}

func NewUserHandler(useCase services.UserUseCase, userMap map[string]models.UserStruct) *UserHandler {
	return &UserHandler{
		UserUseCase: useCase,
		userMap:     userMap,
	}
}

func (ur *UserHandler) UserLogin(c *gin.Context) {
	var UserLoginDetail models.UserLoginRequest

	if err := c.ShouldBindJSON(&UserLoginDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
	}

	err := validator.New().Struct(UserLoginDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return 
	}

	userData, ok := ur.userMap[UserLoginDetail.Email]
	if !ok || userData.Password != UserLoginDetail.Password {
		errs := response.ClientResponse(http.StatusUnauthorized, "Invalid credentials", nil, "invalid")
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	loginDetails, err := ur.UserUseCase.UserLogin(UserLoginDetail.Email, userData.Name)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to user login", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	success := response.ClientResponse(http.StatusCreated, "User successfully login", loginDetails, nil)
	c.JSON(http.StatusCreated, success)
}
