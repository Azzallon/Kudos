package utils

import (
	"kudos/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConvertPointsToMessage(t *testing.T) {
	msg_1 := utils.GetConvertPointsToMessage(-1)
	msg_2 := utils.GetConvertPointsToMessage(0)
	msg_3 := utils.GetConvertPointsToMessage(1)
	msg_4 := utils.GetConvertPointsToMessage(5)
	msg_5 := utils.GetConvertPointsToMessage(10)
	msg_6 := utils.GetConvertPointsToMessage(15)
	msg_7 := utils.GetConvertPointsToMessage(20)
	msg_8 := utils.GetConvertPointsToMessage(25)
	msg_9 := utils.GetConvertPointsToMessage(30)
	msg_10 := utils.GetConvertPointsToMessage(35)
	msg_11 := utils.GetConvertPointsToMessage(40)
	msg_12 := utils.GetConvertPointsToMessage(45)
	msg_13 := utils.GetConvertPointsToMessage(50)
	msg_14 := utils.GetConvertPointsToMessage(55)
	msg_15 := utils.GetConvertPointsToMessage(60)
	msg_16 := utils.GetConvertPointsToMessage(65)
	msg_17 := utils.GetConvertPointsToMessage(70)
	msg_18 := utils.GetConvertPointsToMessage(75)
	msg_19 := utils.GetConvertPointsToMessage(80)
	msg_20 := utils.GetConvertPointsToMessage(85)
	msg_21 := utils.GetConvertPointsToMessage(90)
	msg_22 := utils.GetConvertPointsToMessage(95)
	msg_23 := utils.GetConvertPointsToMessage(100)
	msg_24 := utils.GetConvertPointsToMessage(105)
	msg_25 := utils.GetConvertPointsToMessage(110)
	msg_26 := utils.GetConvertPointsToMessage(115)
	msg_27 := utils.GetConvertPointsToMessage(120)
	msg_28 := utils.GetConvertPointsToMessage(125)
	msg_29 := utils.GetConvertPointsToMessage(130)
	msg_30 := utils.GetConvertPointsToMessage(135)
	msg_31 := utils.GetConvertPointsToMessage(140)
	msg_32 := utils.GetConvertPointsToMessage(145)
	msg_33 := utils.GetConvertPointsToMessage(150)

	expected_1 := "Número de pontos inválido: o número de pontos é <= 0!"
	expected_2 := "Número de pontos inválido: o número de pontos é <= 0!"
	expected_3 := "Número de pontos inválido: o número de pontos informado não é múltiplo de 5!"
	expected_4 := "Você recebeu dois reais em retorno aos kudos OK!"
	expected_5 := "Você recebeu cinco reais em retorno aos kudos NICE!"
	expected_6 := "Você recebeu sete reais em retorno aos kudos NICE, OK!"
	expected_7 := "Você recebeu oito reais em retorno aos kudos GOOD!"
	expected_8 := "Você recebeu dez reais em retorno aos kudos GOOD, OK!"
	expected_9 := "Você recebeu treze reais em retorno aos kudos GOOD, NICE!"
	expected_10 := "Você recebeu quinze reais em retorno aos kudos GOOD, NICE, OK!"
	expected_11 := "Você recebeu dezesseis reais em retorno aos kudos GOOD, GOOD!"
	expected_12 := "Você recebeu dezoito reais em retorno aos kudos GOOD, GOOD, OK!"
	expected_13 := "Você recebeu quinze reais em retorno aos kudos GREAT!"
	expected_14 := "Você recebeu dezessete reais em retorno aos kudos GREAT, OK!"
	expected_15 := "Você recebeu vinte reais em retorno aos kudos GREAT, NICE!"
	expected_16 := "Você recebeu vinte e dois reais em retorno aos kudos GREAT, NICE, OK!"
	expected_17 := "Você recebeu vinte e três reais em retorno aos kudos GREAT, GOOD!"
	expected_18 := "Você recebeu vinte e cinco reais em retorno aos kudos GREAT, GOOD, OK!"
	expected_19 := "Você recebeu vinte e oito reais em retorno aos kudos GREAT, GOOD, NICE!"
	expected_20 := "Você recebeu trinta reais em retorno aos kudos GREAT, GOOD, NICE, OK!"
	expected_21 := "Você recebeu trinta e um reais em retorno aos kudos GREAT, GOOD, GOOD!"
	expected_22 := "Você recebeu trinta e três reais em retorno aos kudos GREAT, GOOD, GOOD, OK!"
	expected_23 := "Você recebeu vinte e cinco reais em retorno aos kudos SUPER!"
	expected_24 := "Você recebeu vinte e sete reais em retorno aos kudos SUPER, OK!"
	expected_25 := "Você recebeu trinta reais em retorno aos kudos SUPER, NICE!"
	expected_26 := "Você recebeu trinta e dois reais em retorno aos kudos SUPER, NICE, OK!"
	expected_27 := "Você recebeu trinta e três reais em retorno aos kudos SUPER, GOOD!"
	expected_28 := "Você recebeu trinta e cinco reais em retorno aos kudos SUPER, GOOD, OK!"
	expected_29 := "Você recebeu trinta e oito reais em retorno aos kudos SUPER, GOOD, NICE!"
	expected_30 := "Você recebeu quarenta reais em retorno aos kudos SUPER, GOOD, NICE, OK!"
	expected_31 := "Você recebeu quarenta e um reais em retorno aos kudos SUPER, GOOD, GOOD!"
	expected_32 := "Você recebeu quarenta e três reais em retorno aos kudos SUPER, GOOD, GOOD, OK!"
	expected_33 := "Você recebeu quarenta reais em retorno aos kudos SUPER, GREAT!"

	assert.Equal(t, expected_1, msg_1)
	assert.Equal(t, expected_2, msg_2)
	assert.Equal(t, expected_3, msg_3)
	assert.Equal(t, expected_4, msg_4)
	assert.Equal(t, expected_5, msg_5)
	assert.Equal(t, expected_6, msg_6)
	assert.Equal(t, expected_7, msg_7)
	assert.Equal(t, expected_8, msg_8)
	assert.Equal(t, expected_9, msg_9)
	assert.Equal(t, expected_10, msg_10)
	assert.Equal(t, expected_11, msg_11)
	assert.Equal(t, expected_12, msg_12)
	assert.Equal(t, expected_13, msg_13)
	assert.Equal(t, expected_14, msg_14)
	assert.Equal(t, expected_15, msg_15)
	assert.Equal(t, expected_16, msg_16)
	assert.Equal(t, expected_17, msg_17)
	assert.Equal(t, expected_18, msg_18)
	assert.Equal(t, expected_19, msg_19)
	assert.Equal(t, expected_20, msg_20)
	assert.Equal(t, expected_21, msg_21)
	assert.Equal(t, expected_22, msg_22)
	assert.Equal(t, expected_23, msg_23)
	assert.Equal(t, expected_24, msg_24)
	assert.Equal(t, expected_25, msg_25)
	assert.Equal(t, expected_26, msg_26)
	assert.Equal(t, expected_27, msg_27)
	assert.Equal(t, expected_28, msg_28)
	assert.Equal(t, expected_29, msg_29)
	assert.Equal(t, expected_30, msg_30)
	assert.Equal(t, expected_31, msg_31)
	assert.Equal(t, expected_32, msg_32)
	assert.Equal(t, expected_33, msg_33)
}
