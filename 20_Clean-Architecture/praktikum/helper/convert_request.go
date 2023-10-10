package helper

import (
	"app/model/domain"
	"app/model/web"
)

func UserCreateRequestToUserDomain(request web.UserCreateRequest) *domain.User {
	return &domain.User{
		Email: request.Email,
		Password: request.Password,
	}
}