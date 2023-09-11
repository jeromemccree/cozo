package generator

import (
	"crypto/rand"
	//"log"

	// "github.com/dchest/authcookie"
	// "github.com/dchest/passwordreset"
	garbler "github.com/michaelbironneau/garbler/lib"

	"encoding/hex"
	//	"github.com/gofrs/uuid"
)

func GenerateValidationCode() (string, error) {
	validationCode, err := garbler.NewPasswords(&garbler.Strong, 1)

	return validationCode[0], err
}

func GenUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7

	return hex.EncodeToString(uuid), err
}

func GenRecruitID() (string, error) {
	uuid := make([]byte, 8)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7

	return hex.EncodeToString(uuid), err
}
