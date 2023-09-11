package owner

import (
	"backend/app/shared/models"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type OwnerAccountRepository struct {
	db *sql.DB
}

func NewOwnerAccountRepository(db *sql.DB) *OwnerAccountRepository {
	return &OwnerAccountRepository{db}
}

// does email already exist
func (Repository *OwnerAccountRepository) DoesOwnerEmailExist(email string) bool {
	_, err := Repository.GetOwnerByEmail(email)

	if err != nil {
		log.Println("Email Already Exist")
		log.Println(err)
		return false
	}

	return true
}

// get owner by email
func (Repository *OwnerAccountRepository) GetOwnerByEmail(email string) (*models.Owner, error) {
	var result models.Owner
	err := Repository.db.QueryRow("select id, public_id, privatekey, usertype, status, firstname, lastname, email, password from owner where email = $1", email).Scan(&result.Id, &result.PublicId, &result.PrivateKey, &result.UserType, &result.Status, &result.Firstname, &result.Lastname, &result.Email, &result.Password)

	if err != nil {
		log.Println(err)

	} else {
		log.Println("Email Already Exist")
	}
	return &result, err

}

func (Repository *OwnerAccountRepository) GetFullOwnerByEmail(email string) (*models.FullProject, error) {
	var result models.FullProject
	err := Repository.db.QueryRow(`SELECT usertype, firstname, lastname, email, password, title, keywords, domain, description, address, city, state, zipcode, project_photo1, project_photo2, project_photo3, project_photo4, project.create_date FROM owner FULL OUTER JOIN project ON project.owner_id = owner.id  WHERE email = $1;`, email).Scan(&result.UserType, &result.Firstname, &result.Lastname, &result.Email, &result.Password, &result.Title, &result.Keywords, &result.Domain, &result.Description, &result.Address, &result.City, &result.State, &result.Zipcode, &result.ProjectPhoto1, &result.ProjectPhoto2, &result.ProjectPhoto3, &result.ProjectPhoto4)

	if err != nil {
		log.Println(err)

	} else {
		log.Println("Email Already Exist")
	}
	return &result, err

}

// create owner
func (Repository *OwnerAccountRepository) AddNewOwner(publicId string, privatekey string, firstname string, lastname string, email string, password string) (*models.Owner, error) {
	var id int
	stmt, err := Repository.db.Prepare(`INSERT INTO owner("public_id", "privatekey", "firstname", "lastname", "email", "password") Values ($1, $2, $3, $4, $5, $6) RETURNING id;`)

	if err != nil {
		return nil, err
	}

	stmtERR := stmt.QueryRow(publicId, privatekey, firstname, lastname, email, password).Scan(&id)

	if stmtERR != nil {
		return nil, stmtERR
	}

	return Repository.GetOwnerById(id)
}

// get owner by id
func (Repository *OwnerAccountRepository) GetOwnerById(id int) (*models.Owner, error) {
	var result models.Owner
	err := Repository.db.QueryRow("select id, public_id, privatekey, usertype, status, firstname,lastname, email, password from owner where id = $1", id).Scan(&result.Id, &result.PublicId, &result.PrivateKey, &result.UserType, &result.Status, &result.Firstname, &result.Lastname, &result.Email, &result.Password)
	if err != nil {

		log.Println(err)
		log.Println("ID")

	}
	return &result, err
}

// get by public_id
func (Repository *OwnerAccountRepository) GetOwnerByPublicId(publicId string) (*models.Owner, error) {
	var result models.Owner
	err := Repository.db.QueryRow("select id, public_id, privatekey, usertype, status, firstname, lastname, email, password from owner where public_id = $1", publicId).Scan(&result.Id, &result.PublicId, &result.PrivateKey, &result.UserType, &result.Status, &result.Firstname, &result.Lastname, &result.Email, &result.Password)
	log.Println(publicId)

	if err != nil {
		log.Println(err)
		log.Println("PublicID Does not exist")
	}
	return &result, err
}

// search
func (Repository *OwnerAccountRepository) GetOwnerBySearch(domain string, zipcode string) (*models.FullProjectResults, error) {

	row, err := Repository.db.Query(`SELECT usertype, firstname, lastname, title, keywords, domain, description, city, state, zipcode, project_photo1, project_photo2, project_photo3, project_photo4, project.create_date FROM owner FULL OUTER JOIN project ON project.owner_id = owner.id WHERE domain = $1 AND zipcode ::text LIKE $2;`, domain, zipcode)
	defer row.Close()

	var results []models.FullProject

	for row.Next() {
		result := models.FullProject{}
		err = row.Scan(&result.UserType, &result.Firstname, &result.Lastname, &result.Title, &result.Keywords, &result.Domain, &result.Description, &result.City, &result.State, &result.Zipcode, &result.ProjectPhoto1, &result.ProjectPhoto2, &result.ProjectPhoto3, &result.ProjectPhoto4, &result.CreateDate)
		results = append(results, result)
	}

	if err != nil {
		log.Println(err)

	}
	Output := &models.FullProjectResults{
		Project: results,
	}

	return Output, err

}

//get All Owner Info
