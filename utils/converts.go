package utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetKudosList(points int) []string {
	values := []int{100, 50, 20, 10, 5}
	i := 0
	max := values[i]
	var kudos []string
	for points > 0 {
		if points >= max {
			points -= max
			kudos = append(kudos, POINTS_TO_KUDOS[max])
		} else {
			i++
			max = values[i]
		}
	}

	return kudos
}

func GetReaisSum(points int) int {
	values := []int{100, 50, 20, 10, 5}
	i := 0
	max := values[i]
	var kudos []string
	var real []int
	for points > 0 {
		if points >= max {
			points -= max
			kudos = append(kudos, POINTS_TO_KUDOS[max])
			real = append(real, KUDOS_TO_REAL[POINTS_TO_KUDOS[max]])
		} else {
			i++
			max = values[i]
		}
	}

	realSum := 0

	for j := 0; j < len(real); j++ {
		realSum += real[j]
	}

	realSumInFill := Convert1to999999(realSum)

	sepKudos := getSepKudos(kudos)

	fmt.Printf("Você recebeu %v reais em retorno aos kudos %v!", realSumInFill, sepKudos)
	return realSum
}

func GetConvertPointsToMessage(points int) string {
	if points > 0 {
		if points%5 == 0 {
			values := []int{100, 50, 20, 10, 5}
			i := 0
			max := values[i]
			var kudos []string
			var real []int
			for points > 0 {
				if points >= max {
					points -= max
					kudos = append(kudos, POINTS_TO_KUDOS[max])
					real = append(real, KUDOS_TO_REAL[POINTS_TO_KUDOS[max]])
				} else {
					i++
					max = values[i]
				}
			}

			realSum := 0

			for j := 0; j < len(real); j++ {
				realSum += real[j]
			}

			realSumInFill := Convert1to999999(realSum)

			sepKudos := getSepKudos(kudos)

			return fmt.Sprintf("Você recebeu %v reais em retorno aos kudos %v!", realSumInFill, sepKudos)
		} else {
			return "Número de pontos inválido: o número de pontos informado não é múltiplo de 5!"
		}
	} else {
		return "Número de pontos inválido: o número de pontos é <= 0!"
	}
}

func ConvertPointsToMessage(c echo.Context) error {
	points, err := strconv.Atoi(c.Param("points"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			"Número de pontos inválido: verifique o tipo do dado passado!")
	}

	if points > 0 {
		if points%5 == 0 {
			values := []int{100, 50, 20, 10, 5}
			i := 0
			max := values[i]
			var kudos []string
			var real []int
			for points > 0 {
				if points >= max {
					points -= max
					kudos = append(kudos, POINTS_TO_KUDOS[max])
					real = append(real, KUDOS_TO_REAL[POINTS_TO_KUDOS[max]])
				} else {
					i++
					max = values[i]
				}
			}

			realSum := 0

			for j := 0; j < len(real); j++ {
				realSum += real[j]
			}

			realSumInFill := Convert1to999999(realSum)

			sepKudos := getSepKudos(kudos)

			msg := fmt.Sprintf(
				"Você recebeu %v reais em retorno aos kudos %v!", realSumInFill, sepKudos)
			return c.JSON(http.StatusOK, msg)
		}

		return c.JSON(http.StatusBadRequest,
			"Número de pontos inválido: o número de pontos informado não é múltiplo de 5!")
	}

	return c.JSON(http.StatusBadRequest, "Número de pontos inválido: o número de pontos é <= 0!")
}

func getSepKudos(kudos []string) string {
	sepKudos := ""

	for k := 0; k < len(kudos); k++ {
		if k != len(kudos)-1 {
			sepKudos += kudos[k] + ", "
		} else {
			sepKudos += kudos[k]
		}
	}

	return sepKudos
}

func ConvertToInt(value string, paramName string) (int, *error) {
	var intId int

	if value != "" {
		var err error

		if intId, err = strconv.Atoi(value); err != nil {
			log.Error(err.Error())
			return 0, &err
		} else {
			return intId, nil
		}
	} else {
		var err error = echo.ErrBadRequest
		return 0, &err
	}
}
