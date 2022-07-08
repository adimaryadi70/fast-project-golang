package model

import "time"

type Authentification struct {
	ID     	   uint  `json:"id" gorm:"primary_key"`
	Username   string `form:"username" gorm: 40; unique json:"username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
}

type SessionToken struct {
	ID            uint      `json:"id" gorm;="primary_key"`
	UserId 	 	  uint      `json:"user_id" binding:"required"`
	Username  	  string    `form:"username" json:"username" binding:"required"`
	StartSession  time.Time `json:"start_session" binding:"required"`
	EndSession    time.Time `json:"end_session" binding:"required"`
}

type UsersInfo struct {
	Token   string `json:"token"`
	Data    interface{}
}