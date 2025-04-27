
package main

import (
	"tll/config"
	"tll/router"
)

func main() {
	config.InitConfig()
	r := router.InitRouterGroup()


	// 启动服务器
	r.Run(config.AppConfig.App.Port) // listen and serve on 0.0.0.0:11451
	
}

