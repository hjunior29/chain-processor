package models

import "gorm.io/gorm"

type ValidProducts struct {
	gorm.Model
	BlockNumber  string
	TimeStamp    string
	Hash         string
	MethodId     string
	FunctionName string
}

type Logs struct {
	gorm.Model
	Error   string
	Message string
}
