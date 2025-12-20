package mail

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gopkg.in/gomail.v2"
)

func SendMail(ctx context.Context, from, to, content string, subject MailSubject) error {
	klog.CtxInfof(ctx, "[Mail Sender] send mail content: %s, from: %s, to: %s", content, from, to)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)

	m.SetHeader("Subject", subject.ToString())
	m.SetBody("text/plain", content)

	if err := mailDial.DialAndSend(m); err != nil {
		klog.CtxErrorf(ctx, "[Mail Sender] SendMail err: %v", err)
		return err
	}

	return nil
}
