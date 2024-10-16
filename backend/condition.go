package main

type ConditionElement struct {
	Int    []string
	String []string
	Bool   []string
}

var BlockCondition = ConditionElement{
	Int:    []string{"x", "y", "z"},
	String: []string{"block_id", "dimension", "player"},
}

var DieCondition = ConditionElement{
	Int: []string{"x", "y", "z"},
	String: []string{
		"dead_id", "dead_name", "killer_id", "killer_name", "dimension",
	},
}

var ChatCondition = ConditionElement{
	String: []string{"player", "message"},
}

var SessionCondition = ConditionElement{
	String: []string{"player"},
	Bool:   []string{"is_join"},
}

var AllowedCondition = map[string]ConditionElement{
	"place":    BlockCondition,
	"break":    BlockCondition,
	"die":      DieCondition,
	"interact": BlockCondition,
	"chat":     ChatCondition,
	"session":  SessionCondition,
}
