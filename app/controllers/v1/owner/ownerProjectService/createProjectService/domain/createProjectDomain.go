package domain

import (
	"backend/app/shared/database"
	"backend/app/shared/models"
	"backend/app/shared/repositories/owner"
	"backend/app/shared/repositories/project"
)

type createProject struct {
}

func NewCreateProject() *createProject {

	return &createProject{}
}

func (Service *createProject) GetOwnerByPublicId(public_id string) (*models.Owner, error) {
	var Repository = owner.NewOwnerAccountRepository(database.BACKENDDB)
	return Repository.GetOwnerByPublicId(public_id)
}

// create owner project
func (Service *createProject) CreateProject(ownerId int, title string, keywords string, domain string, description string, zipcode string, city string, state string, address string, projectPhoto1 string, projectPhoto2 string, projectPhoto3 string, projectPhoto4 string) (*models.Project, error) {
	var Repository = project.NewProjectRepository(database.BACKENDDB)
	return Repository.AddNewProject(ownerId, title, keywords, domain, description, zipcode, city, state, address, projectPhoto1, projectPhoto2, projectPhoto3, projectPhoto4)
}
