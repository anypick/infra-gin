package basegin

import (
	"errors"
	"github.com/anypick/infra"
	"github.com/anypick/infra-gin/helper"
	"github.com/anypick/infra/base/props"
	"github.com/anypick/infra/base/props/container"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"runtime"
	"testing"
)

func InitT() {
	infra.Register(&container.YamlStarter{})
	helper.AddMiddleWares(gin.Recovery(), helper.LogrusMiddle(), helper.AuthMiddleware(true, ParseTokenT))
	Init()

	infra.Register(&infra.BaseInitializerStarter{})
	source := props.NewYamlSource(GetCurrentFilePath("./testx/config.yml", 0))
	app := infra.New(*source)
	app.Start()
}

// 12345678901
// 1234567890
// 12345678
func ParseTokenT(token string) error {
	if len(token) > 10 {
		return errors.New("token is invalid")
	} else if len(token) == 10 {
		return errors.New("token 过期")
	}
	return nil
}

func TestGin(t *testing.T) {
	InitT()
}

// 获取文件的绝对路径
func GetCurrentFilePath(fileName string, skip int) string {
	_, file, _, _ := runtime.Caller(skip)
	// 解析出文件路径
	dir := filepath.Dir(file)
	return filepath.Join(dir, fileName)
}
