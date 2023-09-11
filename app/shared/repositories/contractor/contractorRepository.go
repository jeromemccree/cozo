package contractor

import (
	"backend/app/shared/models"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type ContractorAccountRepository struct {
	db *sql.DB
}

func NewContractorAccountRepository(db *sql.DB) *ContractorAccountRepository {
	return &ContractorAccountRepository{db}
}

// does email already exist
func (Repository *ContractorAccountRepository) DoesContractorEmailExist(email string) bool {
	_, err := Repository.GetContractorByEmail(email)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
	// log.Println("Email Already Exist")

}

func (Repository *ContractorAccountRepository) GetContractorByEmail(email string) (*models.Contractor, error) {
	var result models.Contractor

	err := Repository.db.QueryRow("select email from contractor WHERE email = $1", email).Scan(&result.Email)
	if err != nil {
		log.Println(err)

		// } else {
		// log.Println("Email Already Exist")
	}
	return &result, err

}

func (Repository *ContractorAccountRepository) AddNewContractor(publicId string, privatekey string, userType string, status int, name string, email string, password string) (*models.Contractor, error) {
	var id int
	stmt, err := Repository.db.Prepare(`INSERT INTO contractor("public_id", "privatekey", "usertype", "status", "name", "email", "password")  Values ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`)

	if err != nil {
		return nil, err
	}

	stmtERR := stmt.QueryRow(publicId, privatekey, userType, status, name, email, password).Scan(&id)

	if stmtERR != nil {
		return nil, stmtERR
	}

	return Repository.GetContractorById(id)
}

// get contractor by id
func (Repository *ContractorAccountRepository) GetContractorById(id int) (*models.Contractor, error) {
	var result models.Contractor
	err := Repository.db.QueryRow("select id, public_id, privatekey, usertype, name, email from contractor where id = $1", id).Scan(&result.Id, &result.PublicId, &result.PrivateKey, &result.UserType, &result.Name, &result.Email)
	if err != nil {

		log.Println(err)
		log.Println("ID")

	}
	return &result, err
}

// get by public_id
func (Repository *ContractorAccountRepository) GetContractorByPublicId(publicId string) (*models.Contractor, error) {
	var result models.Contractor
	getPublicId := Repository.db.QueryRow("select id, public_id, usertype, status, name, email from contractor where public_id = $1", publicId).Scan(&result.Id, &result.PublicId, &result.UserType, &result.Status, &result.Name, &result.Email)

	if getPublicId != nil {
		log.Println(getPublicId)

		log.Println("over here")

		log.Println(&result)

	}
	return &result, getPublicId
}
