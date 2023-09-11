package domain

import (
	"backend/app/shared/database"
	"backend/app/shared/models"
	"backend/app/shared/repositories/owner"
	"backend/app/shared/services/validation"
)

type ownerRegistrationService struct {
}

func NewOwnerRegistrationService() *ownerRegistrationService {

	return &ownerRegistrationService{}
}

// get owner by publicId
func (Service *ownerRegistrationService) GetOwnerByPublicId(public_id string) (*models.Owner, error) {
	var Repository = owner.NewOwnerAccountRepository(database.BACKENDDB)
	return Repository.GetOwnerByPublicId(public_id)
}

// is email an actual email
func (Service *ownerRegistrationService) IsValidEmail(email string) bool {
	return validation.ValidateEmail(email)
}

// does owner have a account?
func (Service *ownerRegistrationService) DoesOwnerExist(email string) bool {
	var Repository = owner.NewOwnerAccountRepository(database.BACKENDDB)
	var doesExist = Repository.DoesOwnerEmailExist(email)
	return doesExist
}

// create owner
func (Service *ownerRegistrationService) CreateOwner(public_id string, privatekey string, firstname string, lastname string, email string, password string) (*models.Owner, error) {
	var Repository = owner.NewOwnerAccountRepository(database.BACKENDDB)
	return Repository.AddNewOwner(public_id, privatekey, firstname, lastname, email, password)
}
