package middleware

import (
	"fmt"
	"time"

	"github.com/SuperArnold/GO_Blog/global"
	email "github.com/SuperArnold/GO_Blog/pkg/Email"
	"github.com/SuperArnold/GO_Blog/pkg/app"
	"github.com/SuperArnold/GO_Blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {

	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		IsSSL:    global.EmailSetting.IsSSL,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err: %v"
				global.Logger.WithCallersFrames().Fatalf(s, err)
				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("例外拋錯，發生間： %d", time.Now().Unix()),
					fmt.Sprintf("錯誤訊息： %v", "err"),
				)
				if err != nil {
					global.Logger.Panicf("mail.SendMail err: %v", err)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
