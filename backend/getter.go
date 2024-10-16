package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
	}
}

func FindRecord[T any](c *gin.Context, table string) {

	query := DB.Model(new(T))

	if startTimeRaw := c.Query("start_time"); startTimeRaw != "" {
		startTime, err := time.Parse("2006-01-02 15:04:05", startTimeRaw)
		if err != nil {
			c.JSON(400, Response("你提供的时间格式不正确", err))
			return
		}
		query = query.Where("created_at >= ?", startTime)
	}

	if endTimeRaw := c.Query("end_time"); endTimeRaw != "" {
		endTime, err := time.Parse("2006-01-02 15:04:05", endTimeRaw)
		if err != nil {
			c.JSON(400, Response("你提供的时间格式不正确", err))
			return
		}
		query = query.Where("created_at >= ?", endTime)
	}

	condition := AllowedCondition[table]

	for _, field := range condition.String {
		if value := c.Query(field); value != "" {
			query = query.Where(fmt.Sprintf("%s = ?", field), value)
		}
	}

	for _, field := range condition.Bool {
		if rawValue := c.Query(field); rawValue != "" {
			switch rawValue {
			case "true":
				query = query.Where(fmt.Sprintf("%s = ?", field), true)
			case "false":
				query = query.Where(fmt.Sprintf("%s = ?", field), false)
			default:
				c.JSON(400, Response("布尔类型的参数值不正确", nil))
				return
			}
		}
	}

	for _, field := range condition.Int {
		if rawValue := c.Query(field); rawValue != "" {
			if len(rawValue) < 3 {
				c.JSON(400, Response("字段的值至少需要3字母", nil))
				return
			}
			value, err := strconv.Atoi(rawValue[2:])
			if err != nil {
				c.JSON(400, Response("字段的值应为整数", err))
				return
			}
			switch rawValue[:2] {
			case "lt":
				query = query.Where(fmt.Sprintf("%s < ?", field), value)
			case "gt":
				query = query.Where(fmt.Sprintf("%s > ?", field), value)
			case "eq":
				query = query.Where(fmt.Sprintf("%s = ?", field), value)
			default:
				c.JSON(400, Response("比较符号不正确", err))
				return
			}
		}
	}

	if order := c.Query("order"); order != "" {
		if strings.HasSuffix(order, " asc") ||
			strings.HasSuffix(order, " desc") {
			query = query.Order(order)
		} else {
			c.JSON(400, Response("排序方向不正确", nil))
			return
		}
	}

	limit := 10
	if rawLimit := c.Query("limit"); rawLimit != "" {
		var err error
		if limit, err = strconv.Atoi(rawLimit); err != nil {
			c.JSON(400, Response("限制条数不是数字", err))
			return
		}
		query = query.Limit(limit)
	}

	if rawPage := c.Query("page"); rawPage != "" {
		page, err := strconv.Atoi(rawPage)
		if err != nil {
			c.JSON(200, Response("你提供的页码不是数字", err))
			return
		}
		offset := (page - 1) * limit
		query = query.Offset(offset)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		c.JSON(500, Response("获取记录总数失败", err))
		return
	}

	var data []T
	query.Find(&data)

	c.JSON(200, Response("查询成功", gin.H{
		"page":  count,
		"field": Fields[table],
		"data":  data,
	}))
}
