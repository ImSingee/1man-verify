package user

import (
	"github.com/ImSingee/1man-verify/config"
	"github.com/ImSingee/1man-verify/model"
)

func CreateUser(u *model.User) error {
	result := config.DB.Create(u)
	return result.Error
}
