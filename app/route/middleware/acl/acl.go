package acl

import (
	"backend/app/shared/database"
	"backend/app/shared/models/security/jasonwebtoken"

	"backend/app/shared/repositories/contractor"
	"backend/app/shared/repositories/owner"

	"backend/app/shared/response"

	jwt "github.com/dgrijalva/jwt-go"

	"errors"
	"log"
	"net/http"
	"os"
	"strings"
)

func ForceSSL(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("GO_ENV") == "production" {
			if r.Header.Get("x-forwarded-proto") != "https" {
				sslUrl := "https://" + r.Host + r.RequestURI
				http.Redirect(w, r, sslUrl, http.StatusTemporaryRedirect)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func ForceMyDomain(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("GO_ENV") == "production" {
			var url = r.Host
			if strings.Contains(url, "backend.herokuapp.com") {
				ssUrl := strings.Replace(url, "backend.herokuapp.com", "api.backend.com", -1)
				ssUrl = "https://" + ssUrl + r.RequestURI
				http.Redirect(w, r, ssUrl, http.StatusFound)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func OwnerOnly(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authToken := r.Header.Get("Authorization")
		authArr := strings.Split(authToken, " ")

		if len(authArr) != 2 {
			log.Println("Authentication header is invalid")
			response.SendError(w, http.StatusUnauthorized, "Authentication Header is Invalid.")
			return
		}

		jwtToken := authArr[1]

		claims, err := jwt.ParseWithClaims(jwtToken, &jasonwebtoken.JWTData{}, func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, errors.New("Invalid signing algorithm")
			}
			return []byte(os.Getenv("OWNER_JWT_TOKEN")), nil
		})

		if err != nil {
			log.Println(err)
			response.SendError(w, http.StatusUnauthorized, "Bad Auth Token.")
			return
		}

		data := claims.Claims.(*jasonwebtoken.JWTData)
		OwnerPublic_Id := data.CustomClaims["public_id"]

		var Repository = owner.NewOwnerAccountRepository(database.BACKENDDB)
		Owner, OwnerErr := Repository.GetOwnerByPublicId(OwnerPublic_Id)

		if OwnerErr != nil || Owner.Id == 0 || Owner.PublicId != OwnerPublic_Id {
			response.SendError(w, http.StatusUnauthorized, "Access Denied")
			return
		}

		h.ServeHTTP(w, r)
	})

}

func ContractorOnly(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authToken := r.Header.Get("Authorization")
		authArr := strings.Split(authToken, " ")

		if len(authArr) != 2 {
			log.Println("Authentication header is invalid")
			response.SendError(w, http.StatusUnauthorized, "Authentication Header is Invalid.")
			return
		}

		jwtToken := authArr[1]

		claims, err := jwt.ParseWithClaims(jwtToken, &jasonwebtoken.JWTData{}, func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, errors.New("Invalid signing algorithm")
			}
			return []byte(os.Getenv("WORKER_JWT_TOKEN")), nil
		})

		if err != nil {
			log.Println(err)
			response.SendError(w, http.StatusUnauthorized, "Bad Auth Token.")
			return
		}

		data := claims.Claims.(*jasonwebtoken.JWTData)
		contractorPublic_Id := data.CustomClaims["public_id"]

		var Repository = contractor.NewContractorAccountRepository(database.BACKENDDB)
		contractor, contractorErr := Repository.GetContractorByPublicId(contractorPublic_Id)

		if contractorErr != nil || contractor.Id == 0 || contractor.PublicId != contractorPublic_Id {
			response.SendError(w, http.StatusUnauthorized, "Access Denied")
			return
		}

		h.ServeHTTP(w, r)
	})

}
