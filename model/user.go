package model

import (
	"fmt"
)

type UserInfo struct {
	Uid        int    `json:"uid,omitempty"`
	UserName   string `json:"username"`
	Passwd     string `json:"passwd"`
	Permission string `json:"permission"`
	Email      string `json:"email"`
}

// 看一下有没有必要在model里处理err

func InsertUser(userinfo UserInfo) error {
	_, err := db.Exec("INSERT INTO users(username,passwd,email,permission) values (?,?,?,?)", userinfo.UserName, userinfo.Passwd, userinfo.Email, userinfo.Permission)
	if err != nil {
		return err
	}
	fmt.Println("inserted successfully.")
	return nil
}

func CheckUsername(userinfo UserInfo) (bool, error) {
	var username string
	err := db.Select(&username, "SELECT username FROM users WHERE username=?", userinfo.UserName)
	result := false
	if username != "" {
		result = true
	}
	return result, err
}

// 待debug：用username搜passwd，再进行匹配
func CheckPassword(userinfo UserInfo) (bool, error) {
	var passwd string
	err := db.Select(&passwd, "SELECT passwd FROM users WHERE passwd=?", userinfo.Passwd)
	result := false
	if passwd != "" {
		result = true
	}
	return result, err
}

func DeleteUser(userinfo UserInfo) (string, error) {
	var permission string
	err := db.Select(&permission, "SELECT permission FROM users WHERE username=?", userinfo.UserName)
	if err != nil {
		return "", err
	}
	if permission == "admin" {
		return "you can not delete admin!", nil
	}

	_, err = db.Exec("DELETE FROM users WHERE username=?", userinfo.UserName)
	if err != nil {
		return "delete failed:", err
	}
	return "the user:" + userinfo.UserName + " is deleted successfully", nil
}

func CheckPermission(userinfo UserInfo) (string, error) {
	var permission string
	err := db.Select(&permission, "SELECT permission FROM users WHERE username=?", userinfo.UserName)
	if err != nil {
		return "", err
	}
	return permission, nil
}

func UpdateUserinfo(origUsername string, userinfo UserInfo) (string, error) {
	if userinfo.UserName != "" {
		_, err := db.Exec("UPDATE users SET username=? WHERE username=?", userinfo.UserName, origUsername)
		if err != nil {
			fmt.Println("exec failed, ", err)
			return "update failed", err
		}
	}

	if userinfo.Passwd != "" {
		_, err := db.Exec("UPDATE users SET passwd=? WHERE username=?", userinfo.Passwd, origUsername)
		if err != nil {
			fmt.Println("exec failed, ", err)
			return "update failed", err
		}
	}

	if userinfo.Email != "" {
		_, err := db.Exec("UPDATE users SET email=? WHERE username=?", userinfo.Email, origUsername)
		if err != nil {
			fmt.Println("exec failed, ", err)
			return "update failed", err
		}
	}

	//token中的信息也要改。。。
	if userinfo.Permission=="default" || userinfo.Permission=="admin" {
		_, err := db.Exec("UPDATE users SET permission=? WHERE username=?", userinfo.Permission,origUsername)
		if err != nil {
			fmt.Println("exec failed, ", err)
			return "update failed", err
		}
	}

	return "update success", nil
}

func FindUser(userinfo UserInfo) (UserInfo, error) {
	var info UserInfo
	err := db.Select(&info, "SELECT uid,username,passwd,permission,email FROM users WHERE username=?", userinfo.UserName)
	if err != nil {
		return info, err
	}
	return info, err
}

func ListUsers() ([]string, error) {
	var users []string
	err := db.Select(&users, "SELECT username FROM users")
	if err != nil {
		return users, err
	}
	return users, err
}

// 关闭数据库
