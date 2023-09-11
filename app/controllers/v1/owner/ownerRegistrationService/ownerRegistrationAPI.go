package ownerRegistrationService

import (
	"backend/app/controllers/v1/owner/ownerRegistrationService/domain"
	"backend/app/controllers/v1/owner/ownerRegistrationService/models/request"
	"backend/app/controllers/v1/responseMessages"
	"backend/app/shared/response"

	"backend/app/shared/services/encryption"
	"backend/app/shared/services/generator"

	"backend/app/shared/services/security"
	"log"
	"net/http"

	"encoding/json"
)

func OwnerRegistrationPOST(w http.ResponseWriter, req *http.Request) {

	// json decode
	var ownerRequest request.RegistrationRequest
	decodeErr := json.NewDecoder(req.Body).Decode(&ownerRequest)

	if decodeErr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.CantDecode)
		return
	}

	var ownerregistrationService = domain.NewOwnerRegistrationService()

	// is email a actual email
	isValidEmailAddress := ownerregistrationService.IsValidEmail(ownerRequest.Email)

	if !isValidEmailAddress {
		response.SendJSON(w, responseMessages.InvalidEmailAddress)
		return
	}

	// does owner exist
	var doesownerExist = ownerregistrationService.DoesOwnerExist(ownerRequest.Email)

	if doesownerExist {
		response.SendJSON(w, responseMessages.EmailAlreadyExist)
		return
	}

	// hash password
	hash, hashErr := encryption.HashPassword(ownerRequest.Password)

	if hashErr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	// generate public key
	privatekey, privatekeyErr := generator.GenUUID()

	if privatekeyErr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	// generate public id
	public_id, public_idErr := generator.GenUUID()

	if public_idErr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	// create the owner
	owner, myErr := ownerregistrationService.CreateOwner(public_id, privatekey, ownerRequest.Firstname, ownerRequest.Lastname, ownerRequest.Email, hash)

	if myErr != nil {
		log.Println(myErr)
		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	// create Json Web Token
	JWT := security.GenerateSimpleOwnerJWT(owner.PublicId, owner.UserType, owner.Firstname, owner.Lastname, owner.Email)
	if JWT == "" {
		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	response.SendJSON(w, JWT)

}
