package handlers

import (
	"encoding/json"
	"kudos/domain"
	"kudos/domain/errs"
	"kudos/endpoints/handlers"
	"kudos/tests"
	IMockInterfaces "kudos/tests/mocks/service"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetPeople(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

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

	personService := IMockInterfaces.NewMockIPersonService(controller)
	personService.EXPECT().GetPeople().Return(expected, nil)

	handler := handlers.NewPersonHandlerFromService(personService)

	expectedJson, err := json.Marshal(expected)

	assert.Nil(t, err)
	handlers.RunGenericGetRequestTest(t,
		http.StatusOK,
		handler.GetPeople,
		"/kudos/people",
		tests.NewCtxParams(),
		string(expectedJson),
	)
}

func TestGetPeopleEmpty(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	expected := []domain.Person{}

	personService := IMockInterfaces.NewMockIPersonService(controller)
	personService.EXPECT().GetPeople().Return(expected, nil)

	handler := handlers.NewPersonHandlerFromService(personService)

	expectedJson, err := json.Marshal(expected)

	assert.Nil(t, err)
	handlers.RunGenericGetRequestTest(t,
		http.StatusOK,
		handler.GetPeople,
		"/kudos/people",
		tests.NewCtxParams(),
		string(expectedJson),
	)
}

func TestCreatePerson(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	expected := errs.ApiError{
		StatusCode: 200,
		Message:    "Pessoa cadastrada com sucesso!",
	}

	personService := IMockInterfaces.NewMockIPersonService(controller)
	personService.EXPECT().CreatePerson(gomock.Any()).Return(nil)

	handler := handlers.NewPersonHandlerFromService(personService)

	expectedJson, err := json.Marshal(expected.Error())

	assert.Nil(t, err)
	handlers.RunGenericGetRequestTest(t,
		http.StatusOK,
		handler.CreatePerson,
		"/kudos/person/new",
		tests.NewCtxParams(),
		string(expectedJson),
	)
}

func TestAddPoints(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	expected := errs.ApiError{
		StatusCode: 200,
		Message:    "Os pontos foram adicionados com sucesso!",
	}

	personService := IMockInterfaces.NewMockIPersonService(controller)
	personService.EXPECT().AddPoints(gomock.Any(), gomock.Any()).Return(nil)

	handler := handlers.NewPersonHandlerFromService(personService)

	expectedJson, err := json.Marshal(expected.Error())

	assert.Nil(t, err)
	handlers.RunGenericPutRequestTest(t,
		http.StatusOK,
		handler.AddPoints,
		"/kudos/person/:id/points/add",
		tests.NewCtxParams().Add("id", "1"),
		string(expectedJson),
	)
}

func TestAddPointsErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	expected := errs.ApiError{
		StatusCode: 400,
		Message:    "A pessoa informada não existe!",
	}

	personService := IMockInterfaces.NewMockIPersonService(controller)
	personService.EXPECT().AddPoints(gomock.Any(), gomock.Any()).Return(&expected)

	handler := handlers.NewPersonHandlerFromService(personService)

	expectedJson, err := json.Marshal(expected.Error())

	assert.Nil(t, err)
	handlers.RunGenericGetRequestTest(t,
		http.StatusBadRequest,
		handler.AddPoints,
		"/kudos/person/:id/points/add",
		tests.NewCtxParams().Add("id", "0"),
		string(expectedJson),
	)
}

func TestGetReais(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	expected := 40

	personService := IMockInterfaces.NewMockIPersonService(controller)
	personService.EXPECT().GetReais(gomock.Any()).Return(expected, nil)

	handler := handlers.NewPersonHandlerFromService(personService)

	expectedJson, err := json.Marshal(expected)

	assert.Nil(t, err)
	handlers.RunGenericGetRequestTest(t,
		http.StatusOK,
		handler.GetReais,
		"/kudos/person/:id/reais",
		tests.NewCtxParams().Add("id", "1"),
		string(expectedJson),
	)
}

func TestGetReaisErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	expected := errs.NewBadRequestError("A pessoa informada não existe!")

	personService := IMockInterfaces.NewMockIPersonService(controller)
	personService.EXPECT().GetReais(gomock.Any()).Return(0, expected)

	handler := handlers.NewPersonHandlerFromService(personService)

	expectedJson, err := json.Marshal(expected.Error())

	assert.Nil(t, err)
	handlers.RunGenericGetRequestTest(t,
		http.StatusBadRequest,
		handler.GetReais,
		"/kudos/person/:id/reais",
		tests.NewCtxParams().Add("id", "0"),
		string(expectedJson),
	)
}

func TestGetPointsToMessage(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	expected := "Você recebeu quarenta reais em retorno aos kudos SUPER, GOOD, NICE, OK!"

	personService := IMockInterfaces.NewMockIPersonService(controller)
	personService.EXPECT().GetPointsToMessage(gomock.Any()).Return(expected, nil)

	handler := handlers.NewPersonHandlerFromService(personService)

	expectedJson, err := json.Marshal(expected)

	assert.Nil(t, err)
	handlers.RunGenericGetRequestTest(t,
		http.StatusOK,
		handler.GetPointsToMessage,
		"/kudos/person/:id/message",
		tests.NewCtxParams().Add("id", "1"),
		string(expectedJson),
	)
}

func TestGetPointsToMessageErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	expected := errs.NewBadRequestError("A pessoa informada não existe!")

	personService := IMockInterfaces.NewMockIPersonService(controller)
	personService.EXPECT().GetPointsToMessage(gomock.Any()).Return("", expected)

	handler := handlers.NewPersonHandlerFromService(personService)

	expectedJson, err := json.Marshal(expected.Error())

	assert.Nil(t, err)
	handlers.RunGenericGetRequestTest(t,
		http.StatusBadRequest,
		handler.GetPointsToMessage,
		"/kudos/person/:id/message",
		tests.NewCtxParams().Add("id", "0"),
		string(expectedJson),
	)
}
