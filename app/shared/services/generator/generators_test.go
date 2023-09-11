package generator_test

import (
	"backend/app/shared/services/generator"
	"testing"
)

func ValidationCode(t *testing.T) {

	Code, err := generator.GenerateValidationCode(validationCode)
	if err != nil {
		t.Error("There is no Validationd Code ")
	} else {
		if Code != "" && err == nil {
			t.Errorf("ValidationCode successful")
		}
	}
}

func UUIDCOde(t *testing.T) {

	uniqueId, err := generator.GenrateUniqueId()
	if err != nil {
		t.Error("There is no Unique Id ")
	} else {
		if password != "" && err == nil {
			t.Errorf("UniqueId is Successful")
		}
	}
}
