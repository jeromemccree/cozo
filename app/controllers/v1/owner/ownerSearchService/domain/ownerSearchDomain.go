package domain

import (
	"backend/app/shared/database"
	"backend/app/shared/repositories/contractorProfile"

	"backend/app/shared/models"
)

type ownerSearchService struct {
}

func NewOwnerSearchService() *ownerSearchService {

	return &ownerSearchService{}
}

// search for contractor profile
func (Service *ownerSearchService) SearchForContractor(domain string, zipcode string) (*models.FullContractorResults, error) {
	var Repository = contractorProfile.NewContractorProfileRepository(database.BACKENDDB)
	return Repository.GetContractorProfileBySearch(domain, zipcode)
}
