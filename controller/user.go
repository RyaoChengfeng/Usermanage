package controller

import (
	"Usermanage/config"
	"Usermanage/model"
	"crypto/md5"
	"encoding/hex"
	"github.com/labstack/echo"
	"net/http"
)

func UserRegister(c echo.Context) error {
	username := c.FormValue("username")
	userinfo := model.UserInfo{UserName: username}
	userexit, err := model.CheckUsername(userinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, "")
	}
	if userexit {
		return ErrorHandler(c, http.StatusBadRequest, "username has already exited")
	}

	passwd := c.FormValue("password")
	if passwd == "" {
		return ErrorHandler(c, http.StatusBadRequest, "password is required!")
	}
	h := md5.New()
	h.Write([]byte(passwd))
	encryptedPasswd := hex.EncodeToString(h.Sum([]byte(config.Md5hashSecret)))

	email := c.FormValue("email")
	if email == "" {
		return ErrorHandler(c, http.StatusBadRequest, "email is required!")
	}

	userinfo = model.UserInfo{Passwd: encryptedPasswd, Email: email, Permission: "default"}
	err = model.InsertUser(userinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "register successfully!")
}

func UserLogin(c echo.Context) error {
	username := c.FormValue("username")
	userinfo := model.UserInfo{UserName: username}
	userexit, err := model.CheckUsername(userinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if !userexit {
		return ErrorHandler(c, http.StatusBadRequest, "username is not exited")
	}

	passwd := c.FormValue("password")
	if passwd == "" {
		return ErrorHandler(c, http.StatusBadRequest, "password is required!")
	}
	h := md5.New()
	h.Write([]byte(passwd))
	encryptedPasswd := hex.EncodeToString(h.Sum([]byte(config.Md5hashSecret)))

	userinfo = model.UserInfo{Passwd: encryptedPasswd}
	passwdmach, err := model.CheckPassword(userinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if !passwdmach {
		return ErrorHandler(c, http.StatusBadRequest, "wrong password!")
	}

	userinfo.Permission, err = model.CheckPermission(userinfo)

	t,err := model.CreateToken(userinfo)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
		"massage":"log in successfully",
	})
}

func UserUpdateUserinfo(c echo.Context) error {

}

//func  注销，删除token
func Logout(c echo.Context) error {

}