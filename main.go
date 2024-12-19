package main

import (
	"sales_agent/common/config"
	"sales_agent/common/log"
	"sales_agent/router"
	"sales_agent/task"
)

func main() {
	config.Init()
	log.GetLogger().Info("Hello, Sales Agent!")

	router.SetupRouter().Run(":" + config.GetConf("http.port"))

	task.Init()
}
