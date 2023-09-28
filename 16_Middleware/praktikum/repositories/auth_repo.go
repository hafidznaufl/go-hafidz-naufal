package repositories

import (
	"appmid/config"
	"appmid/model"
)

func Login(password string, email string) (bool, error) {
	var user model.User

	result := config.DB.First(&user, "password = ? AND email = ?", password, email)
	if result.Error != nil {
		return false, result.Error
	}
	
	return true, nil
}