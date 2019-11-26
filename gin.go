package basegin

import (
	"fmt"
	"github.com/anypick/infra"
	"github.com/anypick/infra-gin/config"
	"github.com/anypick/infra-gin/helper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"regexp"
)

var (
	ginEngine *gin.Engine
)

// 对外暴露
func Gin() *gin.Engine {
	return ginEngine
}

type GinStarter struct {
	infra.BaseStarter
}

func (g *GinStarter) Init(ctx infra.StarterContext) {
	ginEngine = initGinApp()
	ginEngine.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})
}

func (g *GinStarter) Start(ctx infra.StarterContext) {
	var (
		engine *gin.Engine
		config = ctx.Yaml()[config.DefaultPrefix].(*config.GinApp)
	)
	engine = Gin()
	routes := engine.Routes()
	helper.IgnorePath = make(map[string]bool)
	go func(routeInfos []gin.RouteInfo) {
		for _, info := range routes {
			flag := false
			logrus.Debugf("API: %s %s %s", info.Method, info.Path, info.Handler)
			for _, ignorePath := range config.AuthIgnore {
				if matched, _ := regexp.Match(ignorePath, []byte(info.Path)); matched {
					helper.IgnorePath[info.Path] = matched
					flag = true
					break
				}
			}
			if !flag {
				helper.IgnorePath[info.Path] = false
			}
		}
		logrus.Debug("auth router match finished")
	}(routes)
	logrus.Debugf("gin start with port %d", config.Port)
	if err := engine.Run(fmt.Sprintf(":%d", config.Port)); err != nil {
		logrus.Error("gin start error, ", err)
	}
}

// web服务是阻塞的
func (g *GinStarter) StartBlocking() bool {
	return true
}

// 初始化gin
func initGinApp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(helper.GetAllMiddleWares()...)
	return app
}
