package main

import (
	"github.com/ImSingee/1man-verify/config"
	"github.com/ImSingee/1man-verify/server"
	"github.com/gin-gonic/gin"
)

func setupServer() (*gin.Engine, error) {

	if !config.Debug() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	server.SetupRouter(r)

	return r, nil
}

func main() {
	engine, err := setupServer()
	if err != nil {
		panic("Failed to setup server")
	}

	err = engine.Run("0.0.0.0:80")
	// TODO 优雅关闭
	if err != nil {
		panic("Failed to run server")
	}
}
