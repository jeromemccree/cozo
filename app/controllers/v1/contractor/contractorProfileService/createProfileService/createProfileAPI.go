package createProfileService

import (
	"backend/app/controllers/v1/contractor/contractorProfileService/createProfileService/domain"
	"backend/app/controllers/v1/contractor/contractorProfileService/models/request"
	"backend/app/controllers/v1/responseMessages"
	"backend/app/shared/response"
	"backend/app/shared/services/security"
	"log"

	"encoding/json"
	"net/http"
)

func CreateProfilePOST(w http.ResponseWriter, r *http.Request) {

	// json decode
	var profileRequest request.ContractorProfileRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&profileRequest)

	if decodeErr != nil {

		log.Println("Cant even decode the json ")
		log.Println(decodeErr)

		response.SendJSON(w, responseMessages.InvalidRequest)
		return
	}

	var profileService = domain.NewCreateProfile()

	//	get contractor by publicId
	contractor, myErr := profileService.GetContractorByPublicId(profileRequest.PublicId)
	if myErr != nil {
		log.Println(myErr)

		log.Println("contractor doesnt have a public id")

		log.Println(contractor)

		response.SendJSON(w, responseMessages.FriendlyError)
		return
	}

	//	does contractor have a profile already
	doesContractorExist := profileService.DoesContractorProfileExist(contractor.Id)
	if doesContractorExist {
		log.Println("contractor has a profile")

		response.SendJSON(w, responseMessages.ContractorAlreadyExist)
		return
	}

	// create the contractor profile
	contractorProfile, myErr := profileService.CreateContractorProfile(contractor.Id, profileRequest.Title, profileRequest.Bio, profileRequest.Domain, profileRequest.Phone, profileRequest.Address, profileRequest.City, profileRequest.State, profileRequest.Zipcode, profileRequest.Url, profileRequest.ProfilePhoto, profileRequest.BackgroundPhoto, profileRequest.TwitterHandle, profileRequest.FacebookHandle, profileRequest.InstagramHandle, profileRequest.LinkedinHandle)

	log.Println("This is contractor profile  ", contractorProfile)

	if contractorProfile == nil || myErr != nil {
		log.Println(myErr)
		log.Println("Problem putting profile into database")

		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)

		return
	}

	// create Json Web Token
	jwt := security.GenerateContractorJWT(contractor.PublicId, contractor.UserType, contractor.Name, contractor.Email, contractorProfile.Title, contractorProfile.Bio, contractorProfile.Domain, contractorProfile.City, contractorProfile.State, contractorProfile.Zipcode, contractorProfile.Url, contractorProfile.ProfilePhoto, contractorProfile.BackgroundPhoto, contractorProfile.Twitterhandle, contractorProfile.Facebookhandle, contractorProfile.Instagramhandle, contractorProfile.Linkedinhandle)

	if jwt == "" {
		log.Println(jwt)

		response.SendError(w, http.StatusBadRequest, responseMessages.JWTProblem)
		return
	}

	response.SendJSON(w, jwt)

}
