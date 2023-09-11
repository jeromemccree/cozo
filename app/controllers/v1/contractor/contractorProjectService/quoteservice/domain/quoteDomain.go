package domain

// import (
// 	"backend/app/shared/database"
// 	"backend/app/shared/models"
// 	"backend/app/shared/repositories/owner"
// 	"backend/app/shared/repositories/project"
// )

// type CreateQuoteService struct {
// }

// func NewCreateQuoteService() *CreateQuoteService {
// 	return &CreateQuoteService{}
// }

// // get owner
// func (Service *CreateQuoteService) GetOwnerByPublicId(public_id string) (*models.Owner, error) {
// 	var Repository = owner.NewOwnerAccountRepository(database.BACKENDDB)
// 	return Repository.GetOwnerByPublicId(public_id)
// }

// // get project
// func (Service *CreateQuoteService) GetProjectId(ownerId int, title string, job string, description string, status string) (*models.Project, error) {
// 	var Repository = project.NewProjectRepository(database.BACKENDDB)
// 	return Repository.GetProjectByProperties(ownerId, title, job, description, status)
// }

// // create quote
// func (Service *CreateQuoteService) CreateQuote(projectId int, employerId int, price int, date string, status string) (*models.Project, error) {
// 	var Repository = project.NewProjectRepository(database.BACKENDDB)
// 	return Repository.CreateQuote(projectId, employerId, price, date, status)
// }

// send email project created
