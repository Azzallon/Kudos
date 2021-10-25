package handlers

import (
	"kudos/data/repository"
	"kudos/endpoints/dto/request"
	"kudos/services"
	"kudos/services/interfaces"
	"kudos/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type PersonHandler struct {
	service interfaces.IPersonService
}

func NewPersonHandler() *PersonHandler {
	r := repository.NewPersonRepositoryDB()
	s := services.NewPersonService(r)
	return &PersonHandler{s}
}

func NewPersonHandlerFromService(service interfaces.IPersonService) *PersonHandler {
	return &PersonHandler{service}
}

func (h PersonHandler) GetPeople(c echo.Context) error {
	person, err := h.service.GetPeople()
	if err != nil {
		return c.String(err.StatusCode, err.Error())
	}
	return c.JSON(http.StatusOK, person)
}

func (h PersonHandler) CreatePerson(c echo.Context) error {
	var person request.Person

	if err := c.Bind(&person); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest,
			"O nome foi passado incorretamente! Verifique o tipo do dado passado!")
	}

	err := h.service.CreatePerson(person.Name)
	if err != nil {
		return c.String(err.StatusCode, err.Message)
	}

	return c.JSON(http.StatusOK, "Pessoa cadastrada com sucesso!")
}

func (h PersonHandler) AddPoints(c echo.Context) error {
	pId := c.Param("id")

	personId, err := utils.ConvertToInt(pId, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "O id informado é inválido!")
	}

	var person request.Person
	if err := c.Bind(&person); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest,
			"Os pontos foram passados incorretamente! Verifique o tipo do dado passado!")
	}

	apiErr := h.service.AddPoints(personId, person.Points)
	if apiErr != nil {
		return c.JSON(http.StatusBadRequest, apiErr.Error())
	}

	return c.JSON(http.StatusOK, "Os pontos foram adicionados com sucesso!")
}

func (h PersonHandler) GetReais(c echo.Context) error {
	pId := c.Param("id")

	personId, err := utils.ConvertToInt(pId, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "O id informado é inválido!")
	}

	reais, apiErr := h.service.GetReais(personId)
	if apiErr != nil {
		return c.JSON(http.StatusBadRequest, apiErr.Error())
	}

	return c.JSON(http.StatusOK, reais)
}

func (h PersonHandler) GetPointsToMessage(c echo.Context) error {
	pId := c.Param("id")

	personId, err := utils.ConvertToInt(pId, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "O id informado é inválido!")
	}

	msg, apiErr := h.service.GetPointsToMessage(personId)
	if apiErr != nil {
		return c.JSON(http.StatusBadRequest, apiErr.Error())
	}

	return c.JSON(http.StatusOK, msg)
}
