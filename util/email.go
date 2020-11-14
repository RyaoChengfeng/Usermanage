package util

import (
	"Usermanage/config"
	"Usermanage/model"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendAuthEmail(userinfo model.UserInfo, tokenString string) (string, error) {
	e := email.Email{
		From:    config.EmailFrom,
		To:      []string{userinfo.Email},
		Subject: "Certification Email",
		HTML:    []byte(`<a href="http://localhost:1323/activate?verify=` + tokenString + `">点击这里验证邮箱</a>`), //
	}
	//e := email.NewEmail()
	//e.From = config.EmailFrom
	//e.To = []string{userinfo.Email} //为啥不能直接用string
	//用tls加密
	fmt.Println("sending email now ...")
	err := e.Send(
		config.EmailAddr,
		smtp.PlainAuth("", config.EmailUsername, config.EmailPasswd, config.EmailHost),
	)
	fmt.Println("finished")
	if err != nil {
		return "", err
	}
	return "Certification Email has been send to your postbox, please enter the email to activate the account", nil
}

//待做：
//发送后要等待一段时间才能发送第二个
//发送第二个后使第一个过期
//一段时间后让邮件过期
