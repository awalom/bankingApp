package main

import (
	"gitlab/awalom/banking/app"
	"gitlab/awalom/banking/logger"
)

func main() {
	logger.InitLogger(app.LogEnv)
	logger.Info("Starting the application")
	app.Start()
}
