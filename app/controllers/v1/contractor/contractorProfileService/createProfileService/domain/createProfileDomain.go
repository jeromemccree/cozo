package domain

import (
	"backend/app/shared/database"
	"backend/app/shared/models"
	"backend/app/shared/repositories/contractor"
	"backend/app/shared/repositories/contractorProfile"
)

type CreateProfile struct {
}

func NewCreateProfile() *CreateProfile {

	return &CreateProfile{}
}

// get contractor by publicId
func (Service *CreateProfile) GetContractorByPublicId(public_id string) (*models.Contractor, error) {
	var Repository = contractor.NewContractorAccountRepository(database.BACKENDDB)
	return Repository.GetContractorByPublicId(public_id)
}

//get contractorProfile by contractor id
func (Service *CreateProfile) DoesContractorProfileExist(contractorId int) bool {
	var Repository = contractorProfile.NewContractorProfileRepository(database.BACKENDDB)
	return Repository.DoesContractorProfileExist(contractorId)
}

// create contractorProfile profile
func (Service *CreateProfile) CreateContractorProfile(contractorId int, title string, bio string, domian string, phone string, address string, city string, state string, zipcode string, url string, profilePhoto string, backgroundPhoto string, twitterHandle string, facebookHandle string, instagramHandle string, linkedinHandle string) (*models.ContractorProfile, error) {
	var Repository = contractorProfile.NewContractorProfileRepository(database.BACKENDDB)
	return Repository.AddContractorProfile(contractorId, title, bio, domian, phone, address, city, state, zipcode, url, profilePhoto, backgroundPhoto, twitterHandle, facebookHandle, instagramHandle, linkedinHandle)
}

// send contractor profile created
