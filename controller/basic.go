package controller

import (
	"Usermanage/config"
	"Usermanage/model"
	"Usermanage/util"
	"crypto/md5"
	"encoding/hex"
	"github.com/labstack/echo"
	"net/http"
)

// /register
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

	token, err := util.CreateUnAuthToken(userinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	msg, err := util.SendAuthEmail(userinfo, token)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	err = model.AddValidToken(token)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token":   token,
		"massage": msg,
	})
}

// /login
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
	userPasswd, err := model.CheckPassword(userinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if encryptedPasswd != userPasswd {
		return ErrorHandler(c, http.StatusBadRequest, "wrong password!")
	}

	activateStatus, err := model.CheckActivateStatus(userinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if activateStatus == 0 {
		return c.JSON(http.StatusForbidden, "this account is not activated")
	}

	userinfo.Permission, err = model.CheckPermission(userinfo)

	token, err := util.CreateAuthToken(userinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	err = model.AddValidToken(token)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token":   token,
		"massage": "log in successfully",
	})
}

// /activate
func UserActivate(c echo.Context) error {
	URLToken := c.QueryParam("verify")
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := util.GetJWTToken(authHeader)
	if URLToken != tokenString {
		return ErrorHandler(c, http.StatusUnauthorized, "username is not match")
	}

	claims, err := util.ParseToken(tokenString)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	username := claims.Username
	userinfo := model.UserInfo{UserName: username}

	msg, err := model.ActivateUser(userinfo)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, msg)
}

//注销
func Logout(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := util.GetJWTToken(authHeader)
	err := model.AddInvalidToken(tokenString)
	if err != nil {
		return ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "logout successfully")
}
