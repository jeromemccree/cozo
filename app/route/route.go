package route

import (

	// registration
	"backend/app/controllers/v1/contractor/contractorRegistrationService"
	"backend/app/controllers/v1/owner/ownerRegistrationService"

	// authentication
	"backend/app/controllers/v1/contractor/contractorAuthenticationService"
	"backend/app/controllers/v1/owner/ownerAuthenticationService"

	// create profile annd project
	"backend/app/controllers/v1/contractor/contractorProfileService/createProfileService"
	"backend/app/controllers/v1/owner/ownerProjectService/createProjectService"

	// search
	"backend/app/controllers/v1/contractor/contractorSearchService"
	"backend/app/controllers/v1/owner/ownerSearchService"

	// "backend/app/route/middleware/acl"
	"backend/app/route/middleware/cors"
	hr "backend/app/route/middleware/httprouterwrapper"
	"backend/app/route/middleware/logrequest"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"

	"github.com/justinas/alice"

	"net/http"
)

func Load() http.Handler {
	return middleware(routes())
}

func routes() *httprouter.Router {
	r := httprouter.New().

	// set 404 handler
	// what was set before
	r.NotFound = alice.New(cors.Handler, acl.ForceMyDomain, acl.ForceSSL).ThenFunc()

	// r.PanicHandler = controllers.ServerError

	// registraton
	r.POST("/contractor/contractorregister", hr.Handler(alice.New(cors.Handler).ThenFunc(contractorRegistrationService.ContractorRegistrationPOST)))
	r.POST("/domain/ownerregister", hr.Handler(alice.New(cors.Handler).ThenFunc(ownerRegistrationService.OwnerRegistrationPOST)))

	// authentiction
	r.POST("/contractor/contractorauth", hr.Handler(alice.New(cors.Handler).ThenFunc(contractorAuthenticationService.ContractorAuthenticationPOST)))
	r.POST("/domain/ownerauth", hr.Handler(alice.New(cors.Handler).ThenFunc(ownerAuthenticationService.OwnerAuthenticationPOST)))

	//create profile and project
	r.POST("/contractor/createprofile", hr.Handler(alice.New(cors.Handler).ThenFunc(createProfileService.CreateProfilePOST)))
	r.POST("/owner/createproject", hr.Handler(alice.New(cors.Handler).ThenFunc(createProjectService.CreateProjectPOST)))

	// search
	r.POST("/contractor/ownersearch", hr.Handler(alice.New(cors.Handler).ThenFunc(contractorSearchService.ContractorSearchPOST)))
	r.POST("/contractorsearch", hr.Handler(alice.New(cors.Handler).ThenFunc(ownerSearchService.OwnerSearchPOST)))

	return r
}

func middleware(h http.Handler) http.Handler {

	//	log everything here
	h = logrequest.Handler(h)

	//Cors for Swagger-ui
	h = cors.Handler(h)

	// clear handler for Gorilla context
	h = context.ClearHandler(h)

	return h
}
