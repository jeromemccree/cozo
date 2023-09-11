package createProjectService

import (
	"backend/app/controllers/v1/owner/ownerProjectService/createProjectService/domain"
	"backend/app/controllers/v1/owner/ownerProjectService/models/request"
	"backend/app/controllers/v1/responseMessages"
	"backend/app/shared/services/security"

	"backend/app/shared/response"
	"log"

	"encoding/json"
	"net/http"
)

func CreateProjectPOST(w http.ResponseWriter, r *http.Request) {

	// json decode
	var projectRequest request.ProjectRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&projectRequest)

	if decodeErr != nil {
		log.Println(decodeErr)
		log.Println("Problem with decode")

		response.SendError(w, http.StatusBadRequest, responseMessages.InvalidRequest)
		return
	}

	var projectDomain = domain.NewCreateProject()

	// get owner
	owner, ownerErr := projectDomain.GetOwnerByPublicId(projectRequest.PublicId)
	if ownerErr != nil {
		log.Println(ownerErr)
		log.Println("Problem with owner getting publicID")

		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	// create project
	project, projectErr := projectDomain.CreateProject(owner.Id, projectRequest.Title, projectRequest.Keywords, projectRequest.Domain, projectRequest.Description, projectRequest.Zipcode, projectRequest.City, projectRequest.State, projectRequest.Address, projectRequest.ProjectPhoto1, projectRequest.ProjectPhoto2, projectRequest.ProjectPhoto3, projectRequest.ProjectPhoto4)

	if project == nil || projectErr != nil {
		log.Println(projectErr)

		log.Println("Problem with creating project")

		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	// create Json Web Token
	jwt := security.GenerateOwnerJWT(owner.UserType, owner.Firstname, owner.Lastname, owner.Email, project.Title, project.Keywords, project.Domain, project.Description, project.Address, project.City, project.State, project.Zipcode, project.ProjectPhoto1, project.ProjectPhoto2, project.ProjectPhoto3, project.ProjectPhoto4)
	if jwt == "" {
		response.SendError(w, http.StatusBadRequest, responseMessages.FriendlyError)
		return
	}

	response.SendJSON(w, jwt)

}
