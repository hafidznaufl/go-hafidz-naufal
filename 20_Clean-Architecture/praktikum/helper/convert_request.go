package helper

import (
	"clean_architecture/model/domain"
	"clean_architecture/model/web"
)

func UserCreateRequestToUserDomain(request web.UserCreateRequest) *domain.User {
	return &domain.User{
		Email: request.Email,
		Password: request.Password,
	}
}