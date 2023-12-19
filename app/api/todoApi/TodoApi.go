package todoApi

import (
	"gologin/app/entity"
	userQo "gologin/app/entity/qo"
	"gologin/app/service/userService"
	//businessError "gologin/error"
	"gologin/resultVo"
	"gologin/utils"
)

func Todo(c *fiber.Ctx) error {
	var loginDto entity.PasswordLoginDto

	// 用户查询
	qo := &userQo.GetUserByMobileQo{Mobile: loginDto.Mobile, MobilePrefix: loginDto.MobilePrefix}
	user := userService.GetUserByMobile(qo)
	token := utils.CreateToken(user.Id, user.PasswordVersion)
	return c.JSON(resultVo.Success(token, c))
}

func T(c *fiber.Ctx) error {
	return c.JSON(resultVo.Success("test", c))
}



