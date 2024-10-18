package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var dialector gorm.Dialector
	var err error
	switch config.DB.Type {
	case "mysql":
		dialector = mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.DB.Username,
			config.DB.Password,
			config.DB.Host,
			config.DB.Port,
			config.DB.DB,
		))
	case "postgres":
		dialector = postgres.Open(fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable Timezone=Asia/Shanghai",
			config.DB.Host,
			config.DB.Username,
			config.DB.Password,
			config.DB.DB,
			config.DB.Port,
		))
	case "sqlite":
		dialector = sqlite.Open(fmt.Sprintf("%s.db", config.DB.DB))
	}
	DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(new(Break))
	err = DB.AutoMigrate(new(Place))
	err = DB.AutoMigrate(new(Interact))
	err = DB.AutoMigrate(new(Die))
	err = DB.AutoMigrate(new(Chat))
	err = DB.AutoMigrate(new(Session))
	err = DB.AutoMigrate(new(Online))
	if err != nil {
		return err
	}

	return nil
}

type CommonStruct struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type BlockStruct struct {
	CommonStruct
	BlockID   string `json:"blockId"`
	Player    string `json:"player"`
	Dimension string `json:"dimension"`
	VectorStruct
}

type VectorStruct struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Break BlockStruct
type Place BlockStruct
type Interact BlockStruct

type Die struct {
	CommonStruct
	DeadID     string `json:"deadId"`
	DeadName   string `json:"deadName"`
	KillerID   string `json:"killerId"`
	KillerName string `json:"killerName"`
	Dimension  string `json:"dimension"`
	VectorStruct
}

type Chat struct {
	CommonStruct
	Player  string `json:"player"`
	Message string `json:"message"`
}

type Session struct {
	CommonStruct
	Player string `json:"player"`
	IsJoin bool   `json:"isJoin"`
}

type Online struct {
	CommonStruct
	Online int    `json:"online"`
	Data   []byte `json:"data"`
}
