package contractorSearchService

import (
	"backend/app/controllers/v1/contractor/contractorSearchService/domain"
	"backend/app/controllers/v1/contractor/contractorSearchService/models/request"
	"backend/app/controllers/v1/responseMessages"
	"backend/app/shared/response"
	"encoding/json"
	"log"
	"net/http"
)

func ContractorSearchPOST(w http.ResponseWriter, req *http.Request) {

	// json decode
	var searchRequest request.SearchRequest
	decodeErr := json.NewDecoder(req.Body).Decode(&searchRequest)

	if decodeErr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.CantDecode)
		log.Println("cant Decode")
		return

	}
	var contractorProfileBySearch = domain.NewContractorSearchService()

	// get search results
	search, searchErr := contractorProfileBySearch.SearchForOwner(searchRequest.Domain, searchRequest.Zipcode)
	if searchErr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.Retry)
		return
	}
	log.Println(search)

	response.SendJSON(w, search)
}
