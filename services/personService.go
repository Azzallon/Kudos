package services

import (
	"kudos/data/interfaces"
	"kudos/domain"
	"kudos/domain/errs"
)

type PersonService struct {
	repo interfaces.IPersonRepository
}

func NewPersonService(repo interfaces.IPersonRepository) *PersonService {
	return &PersonService{repo}
}

func (p PersonService) GetPeople() ([]domain.Person, *errs.ApiError) {
	persons, apiErr := p.repo.GetPeople()
	return persons, apiErr
}

func (p PersonService) CreatePerson(name string) *errs.ApiError {
	err := p.repo.CreatePerson(name)
	return err
}

func (p PersonService) AddPoints(id, points int) *errs.ApiError {
	err := p.repo.AddPoints(id, points)

	if err != nil {
		return err
	}

	return nil
}

func (p PersonService) GetReais(id int) (int, *errs.ApiError) {
	reais, err := p.repo.GetReais(id)

	return reais, err
}

func (p PersonService) GetPointsToMessage(id int) (string, *errs.ApiError) {
	msg, err := p.repo.GetPointsToMessage(id)

	return msg, err
}
