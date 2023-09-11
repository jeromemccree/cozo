package security

import (
	"backend/app/shared/models/security/jasonwebtoken"

	"os"

	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var ownerJwtKey = []byte(os.Getenv("OWNER_JWT_TOKEN"))
var contractorJwtKey = []byte(os.Getenv("WORKER_JWT_TOKEN"))

func GenerateSimpleOwnerJWT(publicId string, usertype string, firstname string, lastname string, email string) string {
	claims := jasonwebtoken.JWTData{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 200).Unix(),
		},
		CustomClaims: map[string]string{
			"publicId":  publicId,
			"usertype":  usertype,
			"firstname": firstname,
			"lastname":  lastname,
			"email":     email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(ownerJwtKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s")
		return "Jwt Error"
	}

	return tokenString
}

func GenerateSimpleContractorJWT(publicId string, usertype string, name string, email string) string {
	claims := jasonwebtoken.JWTData{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 200).Unix(),
		},
		CustomClaims: map[string]string{
			"publicId": publicId,
			"usertype": usertype,
			"name":     name,
			"email":    email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(contractorJwtKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s")
		return "Jwt Error"
	}

	return tokenString
}

func GenerateOwnerJWT(publicId string, usertype string, firstname string, lastname string, email string, title string, keywords string, domain string, description string, city string, state string, zipcode string, projectPhoto1 string, projectPhoto2 string, projectPhoto3 string, projectPhoto4 string) string {
	claims := jasonwebtoken.JWTData{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 200).Unix(),
		},
		CustomClaims: map[string]string{
			"publicId":      publicId,
			"usertype":      usertype,
			"firstname":     firstname,
			"lastname":      lastname,
			"email":         email,
			"title":         title,
			"keywords":      keywords,
			"domain":        domain,
			"description":   description,
			"city":          city,
			"state":         state,
			"zipcode":       zipcode,
			"projectPhoto1": projectPhoto1,
			"projectPhoto2": projectPhoto2,
			"projectPhoto3": projectPhoto3,
			"projectPhoto4": projectPhoto4,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(ownerJwtKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s")
		return "Jwt Error"
	}

	return tokenString
}

func GenerateContractorJWT(publicId string, usertype string, name string, email string, title string, bio string, domain string, city string, state string, zipcode string, url string, profilePhoto string, backgroundPhoto string, twitterHandle string, facebookHandle string, instagramHandle string, linkedinHandle string) string {

	claims := jasonwebtoken.JWTData{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 200).Unix(),
		},
		CustomClaims: map[string]string{
			"publicId":        publicId,
			"usertype":        usertype,
			"name":            name,
			"email":           email,
			"title":           title,
			"bio":             bio,
			"domain":          domain,
			"city":            city,
			"state":           state,
			"zipcode":         zipcode,
			"url":             url,
			"profilePhoto":    profilePhoto,
			"backgroundPhoto": backgroundPhoto,
			"twitterHandle":   twitterHandle,
			"facebookHandle":  facebookHandle,
			"instagramHandle": instagramHandle,
			"linkedinHandle":  linkedinHandle,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(contractorJwtKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s")
		return "Jwt Error"
	}

	return tokenString
}
