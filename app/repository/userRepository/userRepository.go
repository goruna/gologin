package userRepository

import (
	userQo "gologin/app/entity/qo"
	"gologin/global"
	"gologin/model"
)

func GetUserByMobile(qo *userQo.GetUserByMobileQo) *model.UserModel {
	var user = new(model.UserModel)
	tx := global.DB.Where("mobile = ?", qo.Mobile)
	if qo.MobilePrefix == "" {
		qo.MobilePrefix = "86"
	}
	tx.Where("mobile_prefix", qo.MobilePrefix).First(user)
	return user
}

func GetUserById(userId string) *model.UserModel  {
	var user = new(model.UserModel)
	global.DB.Where("id = ?", userId).First(user)
	return user
}
