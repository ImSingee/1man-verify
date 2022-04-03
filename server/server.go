package server

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"

func SetupRouter(r *gin.Engine) {
	Cors(r)

	// TODO: API

	// TODO: Static
}

func Cors(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization", "range")
	r.Use(cors.New(config))
}
