package models

import (
	"time"
)

type Project struct {
	Id            int
	OwnerId       int
	Title         string
	Keywords      string
	Domain        string
	Description   string
	Status        string
	Address       string
	City          string
	State         string
	Zipcode       string
	ProjectPhoto1 string
	ProjectPhoto2 string
	ProjectPhoto3 string
	ProjectPhoto4 string
	CreateDate    time.Time
	UpdateDate    time.Time
}

type FullProjectResults struct {
	Project []FullProject
}

type FullProject struct {
	UserType      string
	Id            int
	PublicId      string
	PrivateKey    string
	Status        int
	Photo         string
	Email         string
	Password      string
	Firstname     string
	Lastname      string
	Title         string
	Keywords      string
	Domain        string
	Description   string
	Address       string
	City          string
	State         string
	Zipcode       string
	ProjectPhoto1 string
	ProjectPhoto2 string
	ProjectPhoto3 string
	ProjectPhoto4 string
	CreateDate    time.Time
	UpdateDate    time.Time
}
