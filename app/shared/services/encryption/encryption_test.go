package encryption_test

import (
	"backend/app/shared/services/encryption"
	"log"
	"testing"
)

func HashPasswordTest(t *testing.T) {
	var password = "Password"

	hashedString, hashErr := encryption.HashPassword(password)

	if hashErr != nil {
		t.Log("hash Test failed")
	} else {
		t.Log("Test Passed")
		t.Log(hashedString)
		log.Println(hashedString)
	}
	return

}
