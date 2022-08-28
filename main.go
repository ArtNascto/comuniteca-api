package main

import (
	"fmt"
	"net/http"
	"time"

	"comuniteca.com.br/api/config"
	"comuniteca.com.br/api/global"
	"comuniteca.com.br/api/handlers"
	"comuniteca.com.br/api/middlewares"
	"comuniteca.com.br/api/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	err := viper.Unmarshal(&global.Config)
	if err != nil {
		utils.Fatal(fmt.Errorf("unable to decode into struct, %v", err))
	}
	config.AssertDatabase()
}

func main() {

	ginMode := gin.DebugMode
	gin.SetMode(ginMode)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           24 * 30 * time.Hour,
	}))

	r.POST("/account/login", handlers.Login)

	api := r.Group("/api")

	api.Use(middlewares.Auth())
	{
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		})
	}
	port := "8080"
	r.Run(":" + port)
}
