package project

import (
	"backend/app/shared/models"
	"database/sql"

	"log"

	_ "github.com/lib/pq"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{db}
}

// get owner by id
func (Repository *ProjectRepository) GetProjectById(id int) (*models.Project, error) {
	var result models.Project
	err := Repository.db.QueryRow("select id, owner_id, title, keywords, domain, description, zipcode, city, state, address, project_photo1, project_photo2, project_photo3, project_photo4 from project where id = $1", id).Scan(&result.Id, &result.OwnerId, &result.Title, &result.Keywords, &result.Domain, &result.Description, &result.Zipcode, &result.City, &result.State, &result.Address, &result.ProjectPhoto1, &result.ProjectPhoto2, &result.ProjectPhoto3, &result.ProjectPhoto4)
	if err != nil {
		log.Println(err)
		log.Println("ID")

	}
	return &result, err
}

// create Project
func (Repository *ProjectRepository) AddNewProject(ownerId int, title string, keywords string, domain string, description string, zipcode string, city string, state string, address string, projectPhoto1 string, projectPhoto2 string, projectPhoto3 string, projectPhoto4 string) (*models.Project, error) {
	var id int

	stmt, err := Repository.db.Prepare(`INSERT INTO project("owner_id", "title", "keywords", "domain", "description",  "zipcode", "city", "state", "address", "project_photo1", "project_photo2", "project_photo3", "project_photo4") Values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);`)

	if err != nil {
		log.Println(err)
		log.Println("Problem with creating the project")

		return nil, err

	}

	stmtERR := stmt.QueryRow(ownerId, title, keywords, domain, description, zipcode, city, state, address, projectPhoto1, projectPhoto2, projectPhoto3, projectPhoto4).Scan(&id)

	if stmtERR != nil {

		return nil, stmtERR
	}

	return Repository.GetProjectById(id)
}

// create Quote
func (Repository *ProjectRepository) CreateQuote(projectId int, employerId int, price int, date string, status string) (*models.Project, error) {
	var id int
	stmt, err := Repository.db.Prepare(`INSERT INTO quote("project_id", "employer_id", "price", "date", "status") Values ($1, $2, $3, $4, $5) RETURNING id;`)

	if err != nil {
		return nil, err
	}

	stmtERR := stmt.QueryRow(projectId, employerId, price, date, status).Scan(&id)

	if stmtERR != nil {
		return nil, stmtERR
	}

	return Repository.CreateQuote(projectId, employerId, price, date, status)
}
