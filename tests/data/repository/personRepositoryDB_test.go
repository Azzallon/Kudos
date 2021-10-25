package repository

import (
	"kudos/data/repository"
	"kudos/domain"
	"kudos/domain/errs"
	"kudos/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPeople(t *testing.T) {
	tests.SetUp(tests.CreatePeople())
	repo := repository.NewPersonRepositoryDB()
	people, err := repo.GetPeople()
	expected := []domain.Person{
		{
			Id:     1,
			Name:   "Felipe",
			Points: 0,
			Reais:  0,
		},
		{
			Id:     2,
			Name:   "kudos",
			Points: 0,
			Reais:  0,
		},
		{
			Id:     3,
			Name:   "Valdinei",
			Points: 0,
			Reais:  0,
		},
		{
			Id:     4,
			Name:   "João",
			Points: 0,
			Reais:  0,
		},
	}
	assert.Equal(t, expected, people)
	assert.Nil(t, err)
}

func TestGetPeopleEmpty(t *testing.T) {
	tests.SetUp(tests.CreatePeopleEmpty())
	repo := repository.NewPersonRepositoryDB()
	people, err := repo.GetPeople()
	expected := []domain.Person{}
	assert.Equal(t, expected, people)
	assert.Nil(t, err)
}

func TestCreatePerson(t *testing.T) {
	repo := repository.NewPersonRepositoryDB()
	err := repo.CreatePerson("José da Silva Santos")
	assert.Nil(t, err)
}

func TestAddPoints(t *testing.T) {
	tests.SetUp(tests.AddPoints())
	repo := repository.NewPersonRepositoryDB()
	err := repo.AddPoints(1, 5)
	assert.Nil(t, err)
}

func TestAddPointsErr(t *testing.T) {
	tests.SetUp(tests.AddPoints())
	repo := repository.NewPersonRepositoryDB()
	err := repo.AddPoints(2, 5)
	expected := errs.NewBadRequestError("A pessoa informada não existe!")
	assert.Equal(t, expected, err)
}

func TestGetReais(t *testing.T) {
	tests.SetUp(tests.GetReais())
	repo := repository.NewPersonRepositoryDB()
	reais, err := repo.GetReais(1)

	assert.Equal(t, 40, reais)
	assert.Nil(t, err)
}

func TestGetReaisErr(t *testing.T) {
	tests.SetUp(tests.GetReais())
	repo := repository.NewPersonRepositoryDB()

	expected := errs.NewBadRequestError("A pessoa informada não existe!")
	reais, err := repo.GetReais(0)

	assert.Equal(t, 0, reais)
	assert.Equal(t, expected, err)
}

func TestGetPointsToMessage(t *testing.T) {
	tests.SetUp(tests.GetPointsToMessage())
	repo := repository.NewPersonRepositoryDB()

	expect := "Você recebeu quarenta reais em retorno aos kudos SUPER, GOOD, NICE, OK!"
	msg, err := repo.GetPointsToMessage(1)

	assert.Equal(t, expect, msg)
	assert.Nil(t, err)
}

func TestGetPointsToMessageErr(t *testing.T) {
	tests.SetUp(tests.GetPointsToMessage())
	repo := repository.NewPersonRepositoryDB()

	expected := errs.NewBadRequestError("A pessoa informada não existe!")
	msg, err := repo.GetPointsToMessage(0)

	assert.Equal(t, "", msg)
	assert.Equal(t, expected, err)
}
