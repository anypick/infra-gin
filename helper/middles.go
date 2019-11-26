package helper

import (
	"github.com/anypick/infra/utils/common"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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

type authMap map[string]interface{}

var IgnorePath map[string]bool

// 鉴权中间件
// 如果negate为true, ignorePath表示需要鉴权的路径，
// 如果negate为false 则表示取反，ignorePath表示不需要鉴权的路径
func AuthMiddleware(negate bool, parseToken func(token string) error) gin.HandlerFunc {
	return func(context *gin.Context) {
		if !IgnorePath[context.Request.URL.Path] && negate {
			token := context.GetHeader("Authorization")
			if common.StrIsBlank(token) {
				context.JSON(http.StatusUnauthorized, authMap{"success": false, "msg": "unauthorized", "total": 0, "rows": nil})
				context.Abort()
				return
			}
			if err := parseToken(token); err != nil {
				context.JSON(http.StatusForbidden, authMap{"success": false,"msg": err.Error(), "total": 0, "rows": nil})
				context.Abort()
				return
			}
		} else {
			context.Next()
		}
	}
}
