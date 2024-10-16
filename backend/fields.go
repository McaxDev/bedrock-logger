package main

type Field struct {
	Key   string
	Title string
}

var Fields = map[string][]Field{
	"break":    BlockField,
	"place":    BlockField,
	"die":      DieField,
	"interact": BlockField,
	"chat":     ChatField,
	"session":  SessionField,
	"online":   OnlineField,
}

var BlockField = []Field{
	{Key: "player", Title: "玩家名"},
	{Key: "block", Title: "方块名"},
	{Key: "dimension", Title: "维度"},
	{Key: "x", Title: "X坐标"},
	{Key: "y", Title: "Y坐标"},
	{Key: "z", Title: "Z坐标"},
}

var DieField = []Field{
	{Key: "dead_id", Title: "死者类型"},
	{Key: "dead_name", Title: "死者名称"},
	{Key: "killer_id", Title: "杀手类型"},
	{Key: "killer_name", Title: "杀手名称"},
	{Key: "dimension", Title: "维度"},
	{Key: "x", Title: "X坐标"},
	{Key: "y", Title: "Y坐标"},
	{Key: "z", Title: "Z坐标"},
}

var ChatField = []Field{
	{Key: "player", Title: "玩家名"},
	{Key: "message", Title: "内容"},
}

var SessionField = []Field{
	{Key: "player", Title: "玩家名"},
	{Key: "is_join", Title: "是否为进服"},
}

var OnlineField = []Field{
	{Key: "data", Title: "在线数据"},
}
