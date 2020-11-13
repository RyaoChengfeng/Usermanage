package model

import (
	"fmt"
)

func AddValidToken(tokenString string) error {
	_, err := cli.Do("Set", tokenString, 1)
	if err != nil {
		return err
	}
	fmt.Println(err)
	return nil
}

func AddInvalidToken(tokenString string) error {
	_, err := cli.Do("Set", tokenString, 0)
	if err != nil {
		return err
	}
	fmt.Println(err)
	return nil
}

//func CheckToken(tokenString string) (int, error) {
//	rst, err := redis.Int(cli.Do("Get", tokenString))
//	if err != nil {
//		return 0, err
//	}
//	return rst, nil
//}
