package contractorAuthenticationService

import (
	"backend/app/controllers/v1/contractor/contractorAuthenticationService/domain"
	"backend/app/controllers/v1/contractor/contractorAuthenticationService/models/request"
	"backend/app/controllers/v1/responseMessages"
	"backend/app/shared/response"
	"backend/app/shared/services/encryption"
	"backend/app/shared/services/security"
	"backend/app/shared/services/validation"
	"log"

	"encoding/json"
	"net/http"
)

func ContractorAuthenticationPOST(w http.ResponseWriter, r *http.Request) {

	//json decode
	var contractorRequest request.AuthenticationRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&contractorRequest)

	if decodeErr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.InvalidRequest)
		return
	}

	// is email a actual email
	isValidEmailAddress := validation.ValidateEmail(contractorRequest.Email)

	if !isValidEmailAddress {
		response.SendError(w, http.StatusBadRequest, responseMessages.InvalidEmailAddress)
		return
	}

	var authenticationService = domain.NewContractorAuthenticationService()

	// does contractor email already exist
	doescontractorEmailExist := authenticationService.DoesContractorEmailExist(contractorRequest.Email)
	if !doescontractorEmailExist {
		response.SendError(w, http.StatusBadRequest, responseMessages.EmailAlreadyExist)
		return
	}

	// does contractor email exist
	contractor, myErr := authenticationService.GetFullContractorByEmail(contractorRequest.Email)
	if myErr != nil {
		log.Println(myErr)
		response.SendError(w, http.StatusBadRequest, responseMessages.CreateContractor)
		return
	}

	// is password equal to password in database
	doescontractorPasswordExist := encryption.CheckPasswordHash(contractor.Password, contractorRequest.Password)
	if !doescontractorPasswordExist {
		response.SendError(w, http.StatusBadRequest, responseMessages.InvalidPassword)
		return
	}

	// create Json Web Token
	jwt := security.GenerateContractorJWT(contractor.PublicId, contractor.UserType, contractor.Name, contractor.Email, contractor.Title, contractor.Bio, contractor.Domain, contractor.City, contractor.State, contractor.Zipcode, contractor.Url, contractor.ProfilePhoto, contractor.BackgroundPhoto, contractor.Twitterhandle, contractor.Facebookhandle, contractor.Instagramhandle, contractor.Linkedinhandle)
	if jwt == "" {
		response.SendError(w, http.StatusBadRequest, responseMessages.JWTProblem)
		return
	}
	response.SendJSON(w, jwt)
}
