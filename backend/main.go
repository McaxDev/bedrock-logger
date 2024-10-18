package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	GetConfig()

	if err := InitDB(); err != nil {
		log.Fatalf("数据初始化失败: %v\n", err)
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.POST("/query/:table", Getter)
	router.GET("/tables", GetTables)
	router.POST("/set", Setter)
	router.Run(fmt.Sprintf(":%s", config.Port))
}
