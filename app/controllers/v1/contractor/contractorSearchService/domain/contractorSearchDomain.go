package domain

import (
	"backend/app/shared/database"
	"backend/app/shared/repositories/owner"

	"backend/app/shared/models"
)

type contractorSearchService struct {
}

func NewContractorSearchService() *contractorSearchService {

	return &contractorSearchService{}
}

// search for onwer listings
func (Service *contractorSearchService) SearchForOwner(domain string, zipcode string) (*models.FullProjectResults, error) {
	var Repository = owner.NewOwnerAccountRepository(database.BACKENDDB)
	return Repository.GetOwnerBySearch(domain, zipcode)
}
