package model

import (
	"fmt"
)

type UserInfo struct {
	Uid        int    `json:"uid,omitempty"` //omitempty用于忽略空字段（这里没有空字段）
	UserName   string `json:"username"`
	Passwd     string `json:"passwd"`
	Permission string `json:"permission"`
	Email      string `json:"email"`
}

func InsertUser(userinfo UserInfo) error {
	fmt.Println("start InsertUser")
	fmt.Println(userinfo)
	_, err := db.Exec("INSERT INTO users(username,passwd,email,permission) values (?,?,?,?)", userinfo.UserName, userinfo.Passwd, userinfo.Email, userinfo.Permission)
	if err != nil {
		return err
	}
	fmt.Println("inserted successfully.")
	return nil
}

func CheckUsername(userinfo UserInfo) (bool, error) {
	fmt.Println("start CheckUsername")
	fmt.Println(userinfo)
	var username []string
	err := db.Select(&username, "SELECT username FROM users WHERE username=?", userinfo.UserName)
	result := false
	if username != nil {
		result = true
		println("username in not exit")
	}
	fmt.Println("finished")
	return result, err
}

func CheckPassword(userinfo UserInfo) (string, error) {
	fmt.Println("start CheckPassword")
	fmt.Println(userinfo)
	var passwd []string
	err := db.Select(&passwd, "SELECT passwd FROM users WHERE username=?", userinfo.UserName)
	fmt.Println("finished")
	return passwd[0], err
}

func CheckActivateStatus(userinfo UserInfo) (int, error) {
	fmt.Println("start CheckActivateStatus")
	var activated []int
	err := db.Select(&activated, "SELECT activated FROM users WHERE username=?", userinfo.UserName)
	fmt.Println("finished")
	return activated[0], err
}

func DeleteUser(userinfo UserInfo) (string, error) {
	fmt.Println("start DeleteUser")
	var permission []string
	err := db.Select(&permission, "SELECT permission FROM users WHERE username=?", userinfo.UserName)
	if err != nil {
		return "", err
	}
	if permission[0] == "admin" {
		return "you can not delete admin!", nil
	}

	_, err = db.Exec("DELETE FROM users WHERE username=?", userinfo.UserName)
	if err != nil {
		return "delete failed:", err
	}
	fmt.Println("finished")
	return "the user:" + userinfo.UserName + " is deleted successfully", nil
}

func CheckPermission(userinfo UserInfo) (string, error) {
	fmt.Println("start CheckPermission")
	var permission []string
	err := db.Select(&permission, "SELECT permission FROM users WHERE username=?", userinfo.UserName)
	if err != nil {
		return "", err
	}
	fmt.Println("finished")
	return permission[0], nil
}

func UpdateUserinfo(origUsername string, userinfo UserInfo) (string, error) {
	fmt.Println("start  UpdateUserinfo")
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

	if userinfo.Permission == "default" || userinfo.Permission == "admin" {
		_, err := db.Exec("UPDATE users SET permission=? WHERE username=?", userinfo.Permission, origUsername)
		if err != nil {
			fmt.Println("exec failed, ", err)
			return "update failed", err
		}
	}

	return "update success", nil
}

func FindUser(userinfo UserInfo) (UserInfo, error) {
	fmt.Println("start FindUser")
	var info []UserInfo
	err := db.Select(&info, "SELECT uid,username,passwd,permission,email FROM users WHERE username=?", userinfo.UserName)
	if err != nil {
		return info[0], err
	}
	return info[0], nil
}

func ListUsers() ([]string, error) {
	fmt.Println("start ListUsers")
	var users []string
	err := db.Select(&users, "SELECT username FROM users")
	if err != nil {
		return users, err
	}
	return users, nil
}

func ActivateUser(userinfo UserInfo) (string, error) {
	fmt.Println("start ActivateUser")
	_, err := db.Exec("UPDATE users SET activated=? WHERE username=?", 1, userinfo.UserName)
	if err != nil {
		return "", err
	}
	return "you account is activate successfully", nil
}
