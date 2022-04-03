package main

import (
	"fmt"
	"github.com/ImSingee/1man-verify/config"
	"github.com/ImSingee/1man-verify/model"
	"github.com/ImSingee/1man-verify/server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.MySQL()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(&model.User{})
	if err != nil {
		return nil, fmt.Errorf("failed on AutoMigrate: %w", err)
	}

	config.DB = db

	return db, nil
}

func setupServer() (*gin.Engine, error) {
	if !config.Debug() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	server.SetupRouter(r)

	return r, nil
}

var cmdServer = &cobra.Command{
	Use: "server",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := setupDB()
		if err != nil {
			panic("Failed  to setup database: " + err.Error())
		}

		engine, err := setupServer()
		if err != nil {
			panic("Failed to setup server: " + err.Error())
		}

		return engine.Run("0.0.0.0:80")
	},
}

func init() {
	app.AddCommand(cmdServer)
}
