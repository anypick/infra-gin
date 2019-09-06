package basegin

import (
	"github.com/anypick/infra"
	"github.com/gin-gonic/gin"
	"log"
)

var ginEngine *gin.Engine

// 对外暴露
func Gin() *gin.Engine {
	return ginEngine
}

type GinStarter struct {
	infra.BaseStarter
}

func (g *GinStarter) Init(ctx infra.StarterContext) {
	ginEngine = initGinApp()
}

func (g *GinStarter) Start(ctx infra.StarterContext) {
	var (
		engine *gin.Engine
		port   string
		e      error
	)
	engine = Gin()
	port = ctx.Yaml().Application.Port;
	routes := engine.Routes()
	for _, info := range routes {
		log.Println(info.Method, info.Path, info.Handler)
	}
	if e = engine.Run(":" + port); e != nil {
		log.Fatal(e)
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
	// 可以处理程序panic，以及500错误
	app.Use(gin.Recovery())
	// 日志
	return app
}
