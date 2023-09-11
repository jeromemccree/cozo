package contractorRegistrationService

import (
	"backend/app/controllers/v1/contractor/contractorRegistrationService/domain"
	"backend/app/controllers/v1/contractor/contractorRegistrationService/models/request"
	"backend/app/controllers/v1/responseMessages"
	"backend/app/shared/response"
	"backend/app/shared/services/security"

	"backend/app/shared/services/encryption"
	"backend/app/shared/services/generator"
	"log"

	"encoding/json"
	"net/http"
)

func ContractorRegistrationPOST(w http.ResponseWriter, r *http.Request) {

	// json decode
	var contractorRequest request.RegistrationRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&contractorRequest)

	if decodeErr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.InvalidRequest)
		return
	}

	var registrationService = domain.NewContractorRegistrationService()

	// is email a actual email
	isValidEmailAddress := registrationService.IsValidEmail(contractorRequest.Email)

	if !isValidEmailAddress {
		response.SendError(w, http.StatusBadRequest, responseMessages.InvalidEmailAddress)

		return
	}
	// does contractor exist already
	var doescontractorExist = registrationService.DoesContractorExist(contractorRequest.Email)

	if doescontractorExist {
		response.SendJSON(w, responseMessages.EmailAlreadyExist)
		return
	}

	// generate public key
	privatekey, privatekeyErr := generator.GenUUID()

	if privatekeyErr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	// generate public id
	publicId, publicIderr := generator.GenUUID()

	if publicIderr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	// Hash Password
	hash, hashErr := encryption.HashPassword(contractorRequest.Password)

	if hashErr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	// create the contractor
	contractor, myErr := registrationService.CreateContractor(publicId, privatekey, contractorRequest.Name, contractorRequest.Email, hash)

	if myErr != nil {
		log.Println(myErr)
		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	// create Json Web Token
	jwt := security.GenerateSimpleContractorJWT(publicId, contractor.UserType, contractor.Name, contractor.Email)
	if jwt == "" {
		response.SendError(w, http.StatusBadRequest, responseMessages.JWTProblem)
		return
	}
	response.SendJSON(w, jwt)
}
