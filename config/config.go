package config

import "time"

const (
 	JWTSecret = "secret"
 	JWTDuration = time.Hour * 12
	//Md5hashSecret = []byte("hash")为什么错了？
 	Md5hashSecret = "hash"
)
