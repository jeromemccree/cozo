package domain

import (
	"backend/app/shared/database"
	"backend/app/shared/models"
	"backend/app/shared/repositories/contractor"

	"backend/app/shared/services/validation"
)

type RegistrationService struct {
}

func NewContractorRegistrationService() *RegistrationService {

	return &RegistrationService{}
}

// get contractor by publicId
func (Service *RegistrationService) GetContractorByPublicId(public_id string) (*models.Contractor, error) {
	var Repository = contractor.NewContractorAccountRepository(database.BACKENDDB)
	return Repository.GetContractorByPublicId(public_id)
}

// is email an actual email
func (Service *RegistrationService) IsValidEmail(email string) bool {
	return validation.ValidateEmail(email)
}

// does contractor have a contractor account?
func (Service *RegistrationService) DoesContractorExist(email string) bool {
	var Repository = contractor.NewContractorAccountRepository(database.BACKENDDB)
	var doesExist = Repository.DoesContractorEmailExist(email)
	return doesExist
}

// create contractor
func (Service *RegistrationService) CreateContractor(public_id string, privatekey string, name string, email string, password string) (*models.Contractor, error) {
	var Repository = contractor.NewContractorAccountRepository(database.BACKENDDB)
	return Repository.AddNewContractor(public_id, privatekey, "contractor", 1, name, email, password)
}

// send welcome email
