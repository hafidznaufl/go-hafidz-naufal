package authenticate

import (
	"appmid/middleware"
	"appmid/model"
	"appmid/repositories"
	"appmid/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LoginController(c echo.Context) error {
	var user model.User
	c.Bind(&user)
	_, err := repositories.Login(user.Email, user.Password)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Email / Password Not Found"))
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Error"))
	}
	tokenResult := middleware.CreateToken(user.Id, user.Name)

	var response model.UserAuthResponse
	response.ID = uint(user.Id)
	response.Email = user.Email
	response.Name = user.Name
	response.Token = tokenResult

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success", response))
}
