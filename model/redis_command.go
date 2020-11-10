package model

type RedisCommandList struct {
	CommandList []*RedisCommand `json:"command_list"  binding:"required"`
}

type RedisCommand struct {
	Method  string                 `json:"method"  binding:"required"`
	Name    string                 `json:"name"  binding:"required"`
	Args    string                 `json:"args"`
	ArgList map[string]interface{} `json:"arg_list"`
}
