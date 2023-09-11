package ownerAuthenticationService

import (
	"backend/app/controllers/v1/owner/ownerAuthenticationService/domain"
	"backend/app/controllers/v1/owner/ownerAuthenticationService/models/request"

	"backend/app/shared/services/encryption"
	"backend/app/shared/services/security"

	"backend/app/controllers/v1/responseMessages"
	"backend/app/shared/response"
	"backend/app/shared/services/validation"
	"log"

	"encoding/json"
	"net/http"
)

func OwnerAuthenticationPOST(w http.ResponseWriter, r *http.Request) {

	//json decode
	var ownerRequest request.AuthenticationRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&ownerRequest)

	if decodeErr != nil {
		response.SendError(w, http.StatusBadRequest, responseMessages.InvalidRequest)
		return
	}
	var authenticationService = domain.

	// is email a actual email
	isValidEmailAddress := validation.ValidateEmail(ownerRequest.Email)

	if !isValidEmailAddress {
		response.SendJSON(w, responseMessages.InvalidEmailAddress)
		return
	}

	//	does owner email exists
	doesownerEmailExist := authenticationService.DoesEmailExist(ownerRequest.Email)
	if !doesownerEmailExist {
		response.SendJSON(w, responseMessages.EmailDoesntExist)
		return
	}

	//	get owner password
	owner, myErr := authenticationService.GetFullOwnerByEmail(ownerRequest.Email)

	if myErr != nil {
		log.Println(myErr)
		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	passwordCorrect := encryption.CheckPasswordHash(owner.Password, ownerRequest.Password)
	if !passwordCorrect {
		response.SendJSON(w, responseMessages.PasswordIncorrect)
		return
	}

	// create Json Web Token
	jwt := security.GenerateOwnerJWT(owner.UserType, owner.Firstname, owner.Lastname, owner.Email, owner.Title, owner.Keywords, owner.Domain, owner.Description, owner.Address, owner.City, owner.State, owner.Zipcode, owner.ProjectPhoto1, owner.ProjectPhoto2, owner.ProjectPhoto3, owner.ProjectPhoto4)
	if jwt == "" {
		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	response.SendJSON(w, jwt)

}
