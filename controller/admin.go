package controller

import (
	"Usermanage/config"
	"Usermanage/model"
	"Usermanage/util"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

// /admin/:user?user=
func AdminDeleteUser(c echo.Context) error {
	username := c.QueryParam("user")
	userinfo := model.UserInfo{UserName: username}
	msg, err := model.DeleteUser(userinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, msg+err.Error())
	}
	return c.JSON(http.StatusOK, msg)
}

// /admin/:user?user=
func AdminFindUserinfo(c echo.Context) error {
	username := c.QueryParam("user")
	userinfo := model.UserInfo{UserName: username}
	info, err := model.FindUser(userinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, info)
}

// /admin/users
func AdminListAllUsers(c echo.Context) error {
	users, err := model.ListUsers()
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// /admin/:user?user=
func AdminUpdateUserinfo(c echo.Context) error {
	origUsername := c.QueryParam("user")
	origUserinfo := model.UserInfo{UserName: origUsername}
	userexit, err := model.CheckUsername(origUserinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if !userexit {
		return ErrorHandler(c, http.StatusBadRequest, "username is not exited")
	}

	uesrname := c.FormValue("NewUsername")
	passwd := c.FormValue("NewPassword")
	email := c.FormValue("NewEmail")
	permission := c.FormValue("permission")
	h := md5.New()
	h.Write([]byte(passwd))
	encryptedPasswd := hex.EncodeToString(h.Sum([]byte(config.Md5hashSecret)))

	var userinfo = model.UserInfo{
		UserName:   uesrname,
		Passwd:     encryptedPasswd,
		Email:      email,
		Permission: permission,
	}

	msg, err := model.UpdateUserinfo(origUsername, userinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, msg)
	}

	if userinfo.Passwd != "" {
		authHeader := c.Request().Header.Get("Authorization")
		tokenString := util.GetJWTToken(authHeader)
		err := model.AddInvalidToken(tokenString)
		if err != nil {
			fmt.Println(err)
			return ErrorHandler(c, http.StatusBadRequest, err.Error())
		}
		msg += ".\n the user's password is changed"
	}
	if userinfo.Permission != "" {
		authHeader := c.Request().Header.Get("Authorization")
		tokenString := util.GetJWTToken(authHeader)
		err := model.AddInvalidToken(tokenString)
		if err != nil {
			fmt.Println(err)
			return ErrorHandler(c, http.StatusBadRequest, err.Error())
		}
		msg += ".\n the user's Permission is changed"
	}

	return c.JSON(http.StatusOK, msg)
}

//admin用什么来操作用户？
//我的构想是进入/admin/:user后填入表单来修改
//还有admin是应该可以改用户的所有东西？包括密码？
