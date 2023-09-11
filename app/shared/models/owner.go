package models

import (
	"time"
)

type Owner struct {
	Id         int
	PublicId   string
	PrivateKey string
	UserType   string
	Status     int
	Firstname  string
	Lastname   string
	Photo      string
	Email      string
	Password   string
	CreateDate time.Time
	UpdateDate time.Time
}
