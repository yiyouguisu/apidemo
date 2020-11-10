package model

type Member struct {
	Id       int    `json:"id"`
	UserName string `json:"username"  binding:"required"`
	NickName string `json:"nickname"  binding:"required"`
	Sex      int    `json:"sex"  binding:"required"`
	Age      int    `json:"age"  binding:"required"`
}
