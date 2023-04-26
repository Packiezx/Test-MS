package model

import "time"

type AccountModel struct {
	Datetime time.Time `bson:"datetime" json:"datetime"`
	Username string    `bson:"username" json:"username"`
	Password string    `bson:"password" json:"password"`
	Name     string    `bosn:"name" json:"name"`
}
