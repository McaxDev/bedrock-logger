package main

import "github.com/gin-gonic/gin"

func GetTables(c *gin.Context) {

	c.JSON(200, Response("查询成功", Tables))
}
