package userService

import (
	userQo "gologin/app/entity/qo"
	"gologin/app/repository/userRepository"
	"gologin/model"
	"gologin/utils"
)

func GetUserByMobile(qo *userQo.GetUserByMobileQo) *model.UserModel {
	return userRepository.GetUserByMobile(qo)
}

func GetPasswordEncrypt(password string, hash string) string  {
	return utils.md5V1(password + hash)
}