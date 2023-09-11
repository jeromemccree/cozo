package domain

import (
	"backend/app/shared/database"
	"backend/app/shared/services/encryption"

	"backend/app/shared/models"
	"backend/app/shared/repositories/owner"
)

type AuthenticationService struct {
}

func NewOwnerAuthenticationService() *AuthenticationService {
	return &AuthenticationService{}
}

func (Service *AuthenticationService) CheckPasswordHash(password string, hash string) bool {
	return encryption.CheckPasswordHash(password, hash)
}

func (Service *AuthenticationService) DoesEmailExist(email string) bool {
	var Repository = owner.NewOwnerAccountRepository(database.BACKENDDB)
	return Repository.DoesOwnerEmailExist(email)
}

func (Service *AuthenticationService) GetFullOwnerByEmail(email string) (*models.FullProject, error) {
	var Repository = owner.NewOwnerAccountRepository(database.BACKENDDB)
	return Repository.GetFullOwnerByEmail(email)
}
