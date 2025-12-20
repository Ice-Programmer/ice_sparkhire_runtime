package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"testing"
)

func TestSendMail(t *testing.T) {
	m := gomail.NewMessage()
	m.SetHeader("From", "2473159069@qq.com")         // 发件人
	m.SetHeader("To", "13706531210@163.com")         // 收件人
	m.SetHeader("Subject", "Go 邮件测试")                // 主题
	m.SetBody("text/html", "<h1>Hello Go Mail</h1>") // 正文
	d := gomail.NewDialer(
		"smtp.qq.com",       // SMTP 服务器
		465,                 // TLS 端口
		"2473159069@qq.com", // 发件邮箱
		"pjkgafhhwvdwdjed",  // SMTP 授权码（不是邮箱密码）
	)

	// 3. 发送
	if err := d.DialAndSend(m); err != nil {
		t.Errorf("SendMail err: %v", err)
	}

	fmt.Println("邮件发送成功")

}
