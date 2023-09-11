package models

import (
	"time"
)

type Contractor struct {
	Id         int
	PublicId   string
	PrivateKey string
	UserType   string
	Status     int
	Name       string
	Email      string
	Password   string
	CreateDate time.Time
	UpdateDate time.Time
}

type FullContractorResults struct {
	ContractorProfile []FullContractorProfile
}

type FullContractorProfile struct {
	PublicId        string
	UserType        string
	Name            string
	Email           string
	Photo           string
	Password        string
	Domain          string
	Title           string
	Bio             string
	City            string
	State           string
	Zipcode         string
	BackgroundPhoto string
	ProfilePhoto    string
	Url             string
	Twitterhandle   string
	Facebookhandle  string
	Instagramhandle string
	Linkedinhandle  string
	CreateDate      time.Time
	UpdateDate      time.Time
}
