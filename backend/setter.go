package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Type string
	Data json.RawMessage
}

func Setter(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if token != "Bearer "+config.Password {
		c.Status(403)
		return
	}

	var body Request
	if err := c.BindJSON(&body); err != nil {
		c.Status(400)
		return
	}

	var err error
	switch body.Type {
	case "break":
		err = CreateRecord[Break](body, c)
	case "place":
		err = CreateRecord[Place](body, c)
	case "interact":
		err = CreateRecord[Interact](body, c)
	case "die":
		err = CreateRecord[Die](body, c)
	case "chat":
		err = CreateRecord[Chat](body, c)
	case "session":
		err = CreateRecord[Session](body, c)
	case "online":
		err = CreateRecord[Online](body, c)
	}
	if err != nil {
		c.Status(500)
		fmt.Println(err)
		return
	}

	c.Status(200)
}

func CreateRecord[T any](body Request, c *gin.Context) error {
	var data T
	if err := json.Unmarshal(
		body.Data, &data,
	); err != nil {
		return err
	}
	return DB.Create(&data).Error
}
