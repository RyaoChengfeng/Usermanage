package controller

import (
	"Usermanage/config"
	"Usermanage/model"
	"crypto/md5"
	"encoding/hex"
	"github.com/labstack/echo"
	"net/http"
)

//要先认证用户是否登录
func UserUpdateUserinfo(c echo.Context) error {
	origUsername := c.Param("user")
	passwd := c.FormValue("NewPassword")
	email := c.FormValue("NewEmail")
	h := md5.New()
	h.Write([]byte(passwd))
	encryptedPasswd := hex.EncodeToString(h.Sum([]byte(config.Md5hashSecret)))

	var userinfo = model.UserInfo{
		Passwd:     encryptedPasswd,
		Email:      email,
	}
	msg, err := model.UpdateUserinfo(origUsername, userinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, msg)
	}

	return c.JSON(http.StatusOK, msg)
}


