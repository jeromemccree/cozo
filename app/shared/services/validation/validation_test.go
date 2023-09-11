package validation_test

import (
	"backend/app/shared/services/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

var validemail = "contractor@building.com"
var invalidemail = "contractor"

func TestValidEmail(t *testing.T) {
	assert := assert.New(t)

	var isValid = validation.ValidateEmail(validemail)

	assert.Equal(true, isValid, "should be valid email address")

}

func TestInvalidValidEmail(t *testing.T) {
	assert := assert.New(t)

	var isValid = validation.ValidateEmail(invalidemail)
	assert.Equal(false, isValid, "Sorry should be valid email address")
}

package passwordvalidation_test
password = "fds.kjnfd"

func TestPasswordValidations(t *testing.T) {
	assert := assert.New(t)

	var isValid = passwordvalidation.VerifyPassword(password)()
	assert.Equal(false, isValid, "Should be equal")

	//	assert.Equal(true, isValid, "Password is valid")
}
