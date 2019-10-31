package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

// 定义一个logrus日志输出的中间件
func LogrusMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		logrus.Infof("%s %d %v %s %s %s",
			"GIN",
			c.Writer.Status(),
			endTime.Sub(startTime),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL)
	}
}
