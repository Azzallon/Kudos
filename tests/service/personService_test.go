package services

import (
	"kudos/domain"
	"kudos/domain/errs"
	"kudos/services"
	IMockInterfaces "kudos/tests/mocks/respository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetPeople(t *testing.T) {
	t.Run("GetPeoplePass", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		persistence := IMockInterfaces.NewMockIPersonRepository(controller)

		expected := []domain.Person{
			{
				Id:     1,
				Name:   "Felipe",
				Points: 0,
				Reais:  0,
			},
			{
				Id:     2,
				Name:   "Jonas",
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
			{
				Id:     5,
				Name:   "Felipe",
				Points: 0,
				Reais:  0,
			},
		}

		persistence.EXPECT().GetPeople().Return(expected, nil)

		service := services.NewPersonService(persistence)

		people, err := service.GetPeople()

		assert.Equal(t, expected, people)
		assert.Nil(t, err)
	})
}

func TestGetPeopleEmpty(t *testing.T) {
	t.Run("GetPeopleEmpty", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		persistence := IMockInterfaces.NewMockIPersonRepository(controller)

		expected := []domain.Person{}

		persistence.EXPECT().GetPeople().Return(expected, nil)

		service := services.NewPersonService(persistence)

		people, err := service.GetPeople()

		assert.Equal(t, expected, people)
		assert.Nil(t, err)
	})
}

func TestCreatePerson(t *testing.T) {
	t.Run("CreatePersonPass", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		persistence := IMockInterfaces.NewMockIPersonRepository(controller)

		expected := errs.ApiError{
			StatusCode: 200,
			Message:    "Pessoa cadastrada com sucesso!",
		}

		persistence.EXPECT().CreatePerson(gomock.Any()).Return(&expected)

		service := services.NewPersonService(persistence)

		err := service.CreatePerson("José da Silva Santos")

		assert.Equal(t, expected.Message, err.Message)
	})
}

func TestAddPoints(t *testing.T) {
	t.Run("AddPointsPass", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		persistence := IMockInterfaces.NewMockIPersonRepository(controller)

		expected := errs.ApiError{
			StatusCode: 200,
			Message:    "Os pontos foram adicionados com sucesso!",
		}

		persistence.EXPECT().AddPoints(gomock.Any(), gomock.Any()).Return(&expected)

		service := services.NewPersonService(persistence)

		err := service.AddPoints(1, 30)

		assert.Equal(t, expected.Message, err.Message)
	})
}

func TestAddPointsErr(t *testing.T) {
	t.Run("AddPointsErr", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		persistence := IMockInterfaces.NewMockIPersonRepository(controller)

		expected := errs.ApiError{
			StatusCode: 400,
			Message:    "A pessoa informada não existe!",
		}

		persistence.EXPECT().AddPoints(gomock.Any(), gomock.Any()).Return(&expected)

		service := services.NewPersonService(persistence)

		err := service.AddPoints(0, 30)

		assert.Equal(t, expected.Message, err.Message)
	})
}

func TestGetReais(t *testing.T) {
	t.Run("GetReaisPass", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		persistence := IMockInterfaces.NewMockIPersonRepository(controller)

		expected := 40

		persistence.EXPECT().GetReais(gomock.Any()).Return(expected, nil)

		service := services.NewPersonService(persistence)

		reais, err := service.GetReais(1)

		assert.Equal(t, expected, reais)
		assert.Nil(t, err)
	})
}

func TestGetReaisErr(t *testing.T) {
	t.Run("GetReaisErr", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		persistence := IMockInterfaces.NewMockIPersonRepository(controller)

		expected := errs.NewBadRequestError("A pessoa informada não existe!")

		persistence.EXPECT().GetReais(gomock.Any()).Return(0, expected)

		service := services.NewPersonService(persistence)

		reais, err := service.GetReais(0)

		assert.Equal(t, 0, reais)
		assert.Equal(t, expected, err)
	})
}

func TestGetPointsToMessage(t *testing.T) {
	t.Run("GetPointsToMessagePass", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		persistence := IMockInterfaces.NewMockIPersonRepository(controller)

		expected := "Você recebeu quarenta reais em retorno aos kudos SUPER, GOOD, NICE, OK!"

		persistence.EXPECT().GetPointsToMessage(gomock.Any()).Return(expected, nil)

		service := services.NewPersonService(persistence)

		msg, err := service.GetPointsToMessage(1)

		assert.Equal(t, expected, msg)
		assert.Nil(t, err)
	})
}

func TestGetPointsToMessageErr(t *testing.T) {
	t.Run("GetPointsToMessageErr", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		persistence := IMockInterfaces.NewMockIPersonRepository(controller)

		expected := errs.NewBadRequestError("A pessoa informada não existe!")

		persistence.EXPECT().GetPointsToMessage(gomock.Any()).Return("", expected)

		service := services.NewPersonService(persistence)

		msg, err := service.GetPointsToMessage(1)

		assert.Equal(t, "", msg)
		assert.Equal(t, expected, err)
	})
}
