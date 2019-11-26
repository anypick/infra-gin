package config

const (
	DefaultPrefix = "application"
)

type GinApp struct {
	Prefix     string
	Port       int      `yaml:"server.port"`
	Name       string   `yaml:"name"`
	AuthIgnore []string `yaml:"auth.ignore"`
}

func (g *GinApp) ConfigAdd(conf map[interface{}]interface{}) {
	g.Port = conf["server.port"].(int)
	g.Name = conf["name"].(string)
	paths := conf["auth.ignore"].([]interface{})
	length := len(paths)
	if length == 0 {
		g.AuthIgnore = make([]string, 2)
		g.AuthIgnore = append(g.AuthIgnore, "/ping")
		g.AuthIgnore = append(g.AuthIgnore, "/favicon.ico")
	} else {
		g.AuthIgnore = make([]string, length+2)
		g.AuthIgnore[0] = "^/ping"
		g.AuthIgnore[1] = "^/favicon.ico"
		for i := 2; i < length+2; i++ {
			g.AuthIgnore[i] = paths[i-2].(string)
		}
	}
}
