package config

import "os"

var (
	EmailFrom = "Ryao <liaocfe@gmail.com>"
	EmailAddr = "smtp.gmail.com:587"
	EmailHost = "smtp.gmail.com"

	EmailUsername = os.Getenv("gmailusername")
	EmailPasswd = os.Getenv("gmailpasswd")
)
