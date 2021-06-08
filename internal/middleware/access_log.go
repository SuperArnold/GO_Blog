package middleware

import (
	"bytes"
	"time"

	"github.com/SuperArnold/GO_Blog/global"
	"github.com/SuperArnold/GO_Blog/pkg/logger"
	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}

	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}

		s := "access log: method: %s , status_code: %d, " + "begin_time: %d, end_time: %d"
		global.Logger.WithFields(fields).Infof(s,
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime)
	}
}

// func Recovery() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				s := "panic recover err: %v"
// 				global.Logger.WithCallersFrames().Fatalf(s, err)
// 				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
// 				c.Abort()
// 			}
// 		}()
// 		c.Next()
// 	}
// }