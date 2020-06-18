package shared

import (
	"fmt"
	"testing"
)

func TestEmailSender_Send(t *testing.T) {
	sender := &EmailSender{
		User:     "flashcard@liuximu.com",
		Password: "Fc295421489",
		Host:     "smtp.mxhichina.com:25",
	}

	fmt.Println("send email")
	err := sender.Send(Email{
		Subject:    "欢迎注册",
		Recipients: []string{"295421489@qq.com"},
		Content: `
<html><body>
	Hello, Dear: </br>
	欢迎使用闪记卡,更好的维护自己的知识体系。
	你的验证码为： </br>
	<h3>999999</h3> </br>
	自动邮件，请勿回复。 </br>

	闪记卡·熙穆
</body></html> `,
		IsHtml: true,
	})

	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
