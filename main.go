package main

import (
	"github.com/mcondero/goportunidades/config"
	"github.com/mcondero/goportunidades/router"
)

var (
	logger config.Logger
)

func main() {
	// Initialize Configs
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("[ERROR] config initialization error: %v", err)
		return
	}
	// Initialize Router
	router.Initialize()
}
