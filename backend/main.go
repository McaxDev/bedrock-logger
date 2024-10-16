package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	GetConfig()

	if err := InitDB(); err != nil {
		log.Fatalf("数据初始化失败: %v\n", err)
	}

	router := gin.Default()
	router.GET("/get", Getter)
	router.POST("/set", Setter)
	router.Run(fmt.Sprintf(":%s", config.Port))
}
