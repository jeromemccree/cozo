package models

import (
	"time"
)

type ContractorProfile struct {
	Id              int
	ContractorId    int
	Title           string
	Domain          string
	Photo           string
	Bio             string
	Phone           string
	Address         string
	City            string
	State           string
	Zipcode         string
	Url             string
	ProfilePhoto    string
	BackgroundPhoto string
	Twitterhandle    string
	Facebookhandle  string
	Instagramhandle string
	Linkedinhandle  string
	CreateDate      time.Time
	UpdateDate      time.Time
}
