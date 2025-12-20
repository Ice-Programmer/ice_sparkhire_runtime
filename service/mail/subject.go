package mail

import "fmt"

type MailSubject int

const (
	VerifyCodeSubject MailSubject = iota
)

func (m MailSubject) ToString() string {
	switch m {
	case VerifyCodeSubject:
		return "发送验证码"
	default:
		return "UNKNOWN"
	}
}

func GenVerifyCodeContent(verifyCode string) string {
	return fmt.Sprintf(
		"尊敬的【用户】：\n"+
			"您好！本次操作的验证码为：\n"+
			"【 %s 】\n"+
			"验证码有效期【有效期：5 分钟】，请尽快完成操作。\n"+
			"如非本人操作，请忽略此邮件。\n"+
			"感谢您使用【SparkHire】！\n"+
			"【Ice Man】团队\n",
		verifyCode,
	)
}
