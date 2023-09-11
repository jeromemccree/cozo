package domain

import (
	"backend/app/shared/database"

	"backend/app/shared/repositories/contractor"
	"backend/app/shared/repositories/contractorProfile"

	"backend/app/shared/models"
)

type AuthenticationService struct {
}

func NewContractorAuthenticationService() *AuthenticationService {
	return &AuthenticationService{}
}

func (Service *AuthenticationService) DoesContractorEmailExist(email string) bool {
	var Repository = contractor.NewContractorAccountRepository(database.BACKENDDB)
	var emailDoesExist = Repository.DoesContractorEmailExist(email)
	return emailDoesExist
}

func (Service *AuthenticationService) GetFullContractorByEmail(email string) (*models.FullContractorProfile, error) {
	var Repository = contractorProfile.NewContractorProfileRepository(database.BACKENDDB)
	return Repository.GetFullContractorByEmail(email)
}
