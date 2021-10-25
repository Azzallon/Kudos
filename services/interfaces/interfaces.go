package interfaces

import (
	"kudos/domain"
	"kudos/domain/errs"
)

type IPersonService interface {
	GetPeople() ([]domain.Person, *errs.ApiError)
	CreatePerson(name string) *errs.ApiError
	AddPoints(id, points int) *errs.ApiError
	GetReais(id int) (int, *errs.ApiError)
	GetPointsToMessage(id int) (string, *errs.ApiError)
}
