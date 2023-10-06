package controller

import (
	"clean_architecture/helper"
	"clean_architecture/model/web"
	"clean_architecture/service"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	CreateUserController() echo.HandlerFunc
	GetAllUserController() echo.HandlerFunc
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (UserController *UserControllerImpl) CreateUserController() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userCreateRequest := web.UserCreateRequest{}
		err := ctx.Bind(&userCreateRequest)
		if err != nil {
			helper.StatusBadRequest(ctx, err)
		}

		response, err := UserController.UserService.CreateUser(userCreateRequest)
		if err != nil {
			if strings.Contains("Email already exist", err.Error()) {
				return helper.StatusBadRequest(ctx, err)
			}

			return helper.StatusInternalServerError(ctx, err)
		}

		userCreateResponse := helper.UserDomainToUserCreateResponse(response)

		token, err := helper.GenerateToken(&userCreateRequest, response.ID)
		if err != nil {
			return helper.StatusInternalServerError(ctx, err)
		}

		userCreateResponse.Token = token

		return helper.StatusCreated(ctx, "Success Create User", userCreateResponse)
	}
}

func (UserController *UserControllerImpl) GetAllUserController() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		response, err := UserController.UserService.GetAllUser()
		if err != nil {
			return helper.StatusInternalServerError(ctx, err)
		}

		return helper.StatusOK(ctx, "Success get all user", response)
	}
}
