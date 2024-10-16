package main

import (
	"encoding/json"
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
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
}

type BlockStruct struct {
	CommonStruct
	BlockID   string
	Player    string
	Dimension string
	X         int
	Y         int
	Z         int
}

type Break BlockStruct
type Place BlockStruct
type Interact BlockStruct

type Die struct {
	CommonStruct
	DeadID     string
	DeadName   string
	KillerID   string
	KillerName string
	Dimension  string
	X          int
	Y          int
	Z          int
}

type Chat struct {
	CommonStruct
	Player  string
	Message string
}

type Session struct {
	CommonStruct
	Player string
	IsJoin bool
}

type Online struct {
	CommonStruct
	Data json.RawMessage
}
