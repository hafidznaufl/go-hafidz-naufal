package helper

import (
	"app/model/domain"
	"app/model/web"
)

func UserDomainToUserCreateResponse(user *domain.User) web.UserCreateResponse {
	return web.UserCreateResponse{
		Email: user.Email,
	}
}