package mail

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

var (
	mailDial *gomail.Dialer
)

func Init(ctx context.Context) error {
	port, err := strconv.ParseInt(os.Getenv("QQ_EMAIL_PORT"), 10, 64)
	if err != nil {
		return fmt.Errorf("QQ_EMAIL_PORT parse error")
	}
	mailDial = gomail.NewDialer(
		os.Getenv("QQ_EMAIL_HOST"),
		int(port),
		os.Getenv("QQ_EMAIL_ADDRESS"),
		os.Getenv("QQ_EMAIL_PASSWORD"),
	)

	klog.CtxInfof(ctx, "init mail dialer successfully!")

	return nil
}
