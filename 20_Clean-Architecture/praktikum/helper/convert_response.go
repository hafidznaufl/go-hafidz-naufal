package helper

import (
	"clean_architecture/model/domain"
	"clean_architecture/model/web"
)

func UserDomainToUserCreateResponse(user *domain.User) web.UserCreateResponse {
	return web.UserCreateResponse{
		Email: user.Email,
	}
}