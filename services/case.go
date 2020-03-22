package services

import (
	"github.com/matgomes/stolen-bike-challenge/models"
	"github.com/matgomes/stolen-bike-challenge/repository"
	"time"
)

type CaseService struct {
	Repo *repository.Repository
}

func (service *CaseService) FindAllCases() ([]models.Case, error) {
	return service.Repo.GetAllCases()
}

func (service *CaseService) FindCase(id string) (models.Case, error) {
	return service.Repo.GetCaseById(id)
}

func (service *CaseService) OpenCase(bike models.Bike) error {

	officer, err := service.findAvailableOfficer()

	if err != nil {
		return err
	}

	c := models.Case{
		Bike:    bike,
		Open:    true,
		Officer: officer,
		Date:    time.Now(),
	}

	if err := service.Repo.InsertCase(c); err != nil {
		return err
	}

	return service.Repo.RemoveOfficer(officer)
}

func (service *CaseService) CloseCase(id string) error {

	foundCase, _ := service.Repo.GetCaseById(id)
	foundCase.Open = false

	err := service.Repo.UpdateCase(foundCase)

	if err != nil {
		return err
	}

	return service.Repo.InsertOfficer(foundCase.Officer)

}

func (service *CaseService) findAvailableOfficer() (models.Officer, error) {
	return service.Repo.FindOfficer()
}
