package contractorProfile

import (
	"backend/app/shared/models"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type ContractorProfileRepository struct {
	db *sql.DB
}

func NewContractorProfileRepository(db *sql.DB) *ContractorProfileRepository {
	return &ContractorProfileRepository{db}
}

// does contractorProfile already exist
func (Repository *ContractorProfileRepository) DoesContractorProfileExist(contractorProfileId int) bool {
	_, err := Repository.GetContractorProfileByContractorId(contractorProfileId)

	if err != nil {

		return false
	}
	return true
}

//  Create Contractor Profile
func (Repository *ContractorProfileRepository) AddContractorProfile(contractorId int, title string, bio string, domian string, phone string, address string, city string, state string, zipcode string, url string, profilePhoto string, backgroundPhoto string, twitterhandle string, facebookhandle string, instagramhandle string, linkedinhandle string) (*models.ContractorProfile, error) {
	var id int
	stmt, err := Repository.db.Prepare(`INSERT INTO contractor_profile("contractor_id", "title", "bio", "domain", "phone", "address", "city", "state", "zipcode", "url", "profile_photo", "background_photo", "twitterhandle", "facebookhandle", "instagramhandle", "linkedinhandle") Values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id;`)

	if err != nil {
		return nil, err
	}

	stmtERR := stmt.QueryRow(contractorId, title, bio, domian, phone, address, city, state, zipcode, url, profilePhoto, backgroundPhoto, twitterhandle, facebookhandle, instagramhandle, linkedinhandle).Scan(&id)
	if stmtERR != nil {
		return nil, stmtERR
	}

	return Repository.GetContractorProfileById(id)

}

// get Contractor by Id
func (Repository *ContractorProfileRepository) GetContractorProfileById(contractorId int) (*models.ContractorProfile, error) {
	var result models.ContractorProfile
	err := Repository.db.QueryRow("select id, contractor_id, title, bio, domain, phone, address, city, state, zipcode, url, profile_photo, background_photo, twitterhandle, facebookhandle, instagramhandle, linkedinhandle from contractor_profile where id = $1", contractorId).Scan(&result.Id, &result.ContractorId, &result.Title, &result.Bio, &result.Domain, &result.Phone, &result.Address, &result.City, &result.State, &result.Zipcode, &result.Url, &result.ProfilePhoto, &result.BackgroundPhoto, &result.Twitterhandle, &result.Facebookhandle, &result.Instagramhandle, &result.Linkedinhandle)
	if err != nil {

		log.Println(err)

	}
	return &result, err
}

// get contractorProfile by contractorProfile Id
func (Repository *ContractorProfileRepository) GetContractorProfileByContractorId(contractorProfileId int) (*models.ContractorProfile, error) {
	var result models.ContractorProfile
	err := Repository.db.QueryRow("select id, contractor_id, title, bio, domain, phone, address, city, state, zipcode, url, profile_photo, background_photo, twitterhandle, facebookhandle, instagramhandle, linkedinhandle from contractor_profile where contractor_id = $1", contractorProfileId).Scan(&result.Id, &result.ContractorId, &result.Title, &result.Bio, &result.Domain, &result.Phone, &result.Address, &result.City, &result.State, &result.Zipcode, &result.Url, &result.ProfilePhoto, &result.BackgroundPhoto, &result.Twitterhandle, &result.Facebookhandle, &result.Instagramhandle, &result.Linkedinhandle)

	if err != nil {
		log.Println(err)
		log.Println("ID")

	}
	return &result, err
}

// get full contractor by email
func (Repository *ContractorProfileRepository) GetFullContractorByEmail(email string) (*models.FullContractorProfile, error) {
	var result models.FullContractorProfile
	err := Repository.db.QueryRow(`SELECT public_id, usertype, name, password, title, bio, profile_photo, background_photo, domain, city, state, zipcode, url, twitterhandle, facebookhandle, instagramhandle, linkedinhandle, contractor.create_date FROM contractor FULL OUTER JOIN contractor_profile ON contractor.id = contractor_profile.contractor_id WHERE email = $1;`, email).Scan(&result.UserType, &result.Name, &result.Password, &result.Title, &result.Bio, &result.ProfilePhoto, &result.BackgroundPhoto, &result.Domain, &result.City, &result.State, &result.Zipcode, &result.Url, &result.Twitterhandle, &result.Facebookhandle, &result.Instagramhandle, &result.Linkedinhandle, &result.CreateDate)

	if err != nil {
		log.Println(err)
		log.Println("ID")

	}
	return &result, err
}

// search
func (Repository *ContractorProfileRepository) GetContractorProfileBySearch(domain string, zipcode string) (*models.FullContractorResults, error) {
	row, err := Repository.db.Query(`SELECT usertype, name, title, bio, profile_photo, background_photo, domain, city, state, zipcode, url, twitterhandle, facebookhandle, instagramhandle, linkedinhandle FROM contractor FULL OUTER JOIN contractor_profile ON contractor.id = contractor_profile.contractor_id WHERE domain = $1 AND zipcode ::text LIKE $2;`, domain, zipcode)
	defer row.Close()

	var results []models.FullContractorProfile

	for row.Next() {
		result := models.FullContractorProfile{}
		err = row.Scan(&result.UserType, &result.Name, &result.Title, &result.Bio, &result.ProfilePhoto, &result.BackgroundPhoto, &result.Domain, &result.City, &result.State, &result.Zipcode, &result.Url, &result.Twitterhandle, &result.Facebookhandle, &result.Instagramhandle, &result.Linkedinhandle)
		results = append(results, result)

	}
	if err != nil {
		log.Println(err)

	}
	Output := &models.FullContractorResults{
		ContractorProfile: results,
	}

	return Output, err

}
