package ownerSearchService

import (
	"backend/app/controllers/v1/owner/ownerSearchService/domain"
	"backend/app/controllers/v1/owner/ownerSearchService/models/request"
	"backend/app/controllers/v1/responseMessages"
	"backend/app/shared/response"
	"encoding/json"
	"log"
	"net/http"
)

func OwnerSearchPOST(w http.ResponseWriter, req *http.Request) {
	log.Println("Hello")

	// json decode
	var searchRequest request.SearchRequest
	decodeErr := json.NewDecoder(req.Body).Decode(&searchRequest)

	if decodeErr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.CantDecode)
		log.Println("cant Decode")
		return

	}
	var ownerProfileBySearch = domain.NewOwnerSearchService()

	// get search results
	search, searchErr := ownerProfileBySearch.SearchForContractor(searchRequest.Domain, searchRequest.Zipcode)
	log.Println("Search Error")

	if searchErr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.Retry)
		return
	}
	response.SendJSON(w, search)
}
