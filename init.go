package basegin

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra-gin/config"
	"github.com/anypick/infra/base/props/container"
)

// 用户调用,初始化资源
func Init() {
	container.RegisterYamContainer(&config.GinApp{Prefix: config.DefaultPrefix})
	infra.Register(&GinStarter{})
}
