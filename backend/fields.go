package main

type Table struct {
	Name    string
	Display string
	Fields  []Field
}

type Field struct {
	Key   string
	Title string
	Type  string
}

var Tables = []Table{
	{Name: "break", Display: "方块破坏", Fields: BlockField},
	{Name: "place", Display: "方块放置", Fields: BlockField},
	{Name: "interact", Display: "方块互动", Fields: BlockField},
	{Name: "die", Display: "实体死亡", Fields: DieField},
	{Name: "chat", Display: "玩家聊天", Fields: ChatField},
	{Name: "session", Display: "玩家进退", Fields: SessionField},
	{Name: "online", Display: "玩家在线", Fields: OnlineField},
}

var BlockField = []Field{
	{Key: "player", Title: "玩家名", Type: "string"},
	{Key: "blockId", Title: "方块名", Type: "string"}, 
	{Key: "dimension", Title: "维度", Type: "string"},
	{Key: "x", Title: "X坐标", Type: "number"},
	{Key: "y", Title: "Y坐标", Type: "number"},
	{Key: "z", Title: "Z坐标", Type: "number"},
}

var DieField = []Field{
	{Key: "deadId", Title: "死者类型", Type: "string"},
	{Key: "deadName", Title: "死者名称", Type: "string"},
	{Key: "killerId", Title: "杀手类型", Type: "string"},
	{Key: "killerName", Title: "杀手名称", Type: "string"},
	{Key: "dimension", Title: "维度", Type: "string"},
	{Key: "x", Title: "X坐标", Type: "number"},
	{Key: "y", Title: "Y坐标", Type: "number"},
	{Key: "z", Title: "Z坐标", Type: "number"},
}

var ChatField = []Field{
	{Key: "player", Title: "玩家名", Type: "string"},
	{Key: "message", Title: "内容", Type: "string"},
}

var SessionField = []Field{
	{Key: "player", Title: "玩家名", Type: "string"},
	{Key: "isJoin", Title: "是否为进服", Type: "boolean"},
}

var OnlineField = []Field{
	{Key: "data", Title: "在线数据", Type: "null"},
}
