package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Getter(c *gin.Context) {

	table := c.Param("table")

	switch table {
	case "break":
		FindRecord[Break](c, table)
	case "place":
		FindRecord[Place](c, table)
	case "die":
		FindRecord[Die](c, table)
	case "interact":
		FindRecord[Interact](c, table)
	case "chat":
		FindRecord[Chat](c, table)
	case "session":
		FindRecord[Session](c, table)
	case "online":
		FindRecord[Online](c, table)
	}
}

type QueryRequest struct {
	StartTime string
	EndTime   string
	Order     string
	isAsc     bool
	Limit     int
	Page      int
	StartX    int
	EndX      int
	StartY    int
	EndY      int
	StartZ    int
	EndZ      int
	BlockID   string
	Player    string
	Dimension string
	IsJoin    bool
}

func FindRecord[T any](c *gin.Context, tableName string) {

	var request QueryRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, Response("无法读取你的请求", err))
		return
	}

	query := DB.Model(new(T))

	// 起始时间 startTime
	if rawStartTime := request.StartTime; rawStartTime != "" {
		startTime, err := time.Parse("2006-01-02 15:04:05", rawStartTime)
		if err != nil {
			c.JSON(400, Response("你提供的时间格式不正确", err))
			return
		}
		query = query.Where("created_at >= ?", startTime)
	}

	// 结束时间 endTime
	if rawEndTime := request.EndTime; rawEndTime != "" {
		endTime, err := time.Parse("2006-01-02 15:04:05", rawEndTime)
		if err != nil {
			c.JSON(400, Response("你提供的时间格式不正确", err))
			return
		}
		query = query.Where("created_at >= ?", endTime)
	}

	// 记录排序顺序 order
	if order := request.Order; order != "" {
		sequence := "desc"
		if request.isAsc {
			sequence = "asc"
		}
		query = query.Order(fmt.Sprintf("%s %s", order, sequence))
	}

	var table Table
	for _, value := range Tables {
		if tableName == value.Name {
			table = value
		}
	}

	// 特定查询条件
	for _, field := range table.Fields {
		switch field.Key {
		case "blockId":
			query.Scopes(StringScope("block_id", request.BlockID))
		case "player":
			query.Scopes(StringScope("player", request.Player))
		case "dimension":
			query.Scopes(StringScope("dimension", request.Dimension))
		case "startX":
			query.Scopes(IntScope("x", request.StartX, true))
		case "endX":
			query.Scopes(IntScope("x", request.EndX, false))
		case "startY":
			query.Scopes(IntScope("y", request.StartY, true))
		case "endY":
			query.Scopes(IntScope("y", request.EndY, false))
		case "startZ":
			query.Scopes(IntScope("z", request.StartZ, true))
		case "endZ":
			query.Scopes(IntScope("z", request.EndZ, false))
		case "jsJoin":
			query.Scopes(BoolScope("is_join", request.IsJoin))
		}
	}

	// 获取记录数
	var count int64
	if err := query.Count(&count).Error; err != nil {
		c.JSON(500, Response("获取记录总数失败", err))
		return
	}

	// 限制查询条目 limit
	limit := 10
	if request.Limit != 0 {
		limit = request.Limit
	}
	query = query.Limit(limit)

	// 查询页数 page
	if request.Page != 0 {
		query = query.Offset((request.Page - 1) * limit)
	}

	// 查询数据并返回响应
	var data []T
	query.Find(&data)
	c.JSON(200, Response("查询成功", gin.H{
		"count": count,
		"field": table.Fields,
		"data":  data,
	}))
}

type GormScope func(*gorm.DB) *gorm.DB

func StringScope(condition, value string) GormScope {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			db = db.Where(fmt.Sprintf("%s = ?", condition), value)
		}
		return db
	}
}

func IntScope(condition string, value int, isStart bool) GormScope {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			sign := "<="
			if isStart {
				sign = ">="
			}
			db.Where(fmt.Sprintf("%s %s ?", condition, sign), value)
		}
		return db
	}
}

func BoolScope(condition string, value bool) GormScope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s = ?", condition), value)
	}
}
