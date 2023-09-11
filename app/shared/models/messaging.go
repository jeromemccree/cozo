package models

import (
	"time"
)

type Chat struct {
	Id         int
	User1Id    int
	User2Id    int
	CreateDate time.Time
	UpdateDate time.Time
}

type Message struct {
	Id          int
	ChatId      int
	SentFrom    int
	SentTo      int
	Message     string
	MessageType string
	Attatchment string
	CreateDate  time.Time
	UpdateDate  time.Time
}
