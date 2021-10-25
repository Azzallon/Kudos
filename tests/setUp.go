package tests

import (
	"kudos/data/db"
)

func SetUp(queries []string) {
	db := db.GetConnection()

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			panic(err)
		}
	}

	defer db.Close()
}

func CreatePeople() []string {
	commands := []string{
		`DELETE FROM kudos WHERE id != 0`,
		`DELETE FROM person WHERE id != 0`,
		`ALTER TABLE kudos AUTO_INCREMENT = 1`,
		`ALTER TABLE person AUTO_INCREMENT = 1`,
		`INSERT INTO person (id, name, points, reais) VALUES 
			(DEFAULT, 'Felipe', DEFAULT, DEFAULT), 
			(DEFAULT, 'kudos', DEFAULT, DEFAULT),
			(DEFAULT, 'Valdinei', DEFAULT, DEFAULT),
			(DEFAULT, 'João', DEFAULT, DEFAULT)`,
	}

	return commands
}

func CreatePeopleEmpty() []string {
	commands := []string{
		`DELETE FROM kudos WHERE id != 0`,
		`DELETE FROM person WHERE id != 0`,
		`ALTER TABLE kudos AUTO_INCREMENT = 1`,
		`ALTER TABLE person AUTO_INCREMENT = 1`,
	}

	return commands
}

func AddPoints() []string {
	commands := []string{
		`DELETE FROM kudos WHERE id != 0`,
		`DELETE FROM person WHERE id != 0`,
		`ALTER TABLE kudos AUTO_INCREMENT = 1`,
		`ALTER TABLE person AUTO_INCREMENT = 1`,
		`INSERT INTO person (id, name, points, reais) 
			VALUES (DEFAULT, 'José da Silva Santos', DEFAULT, DEFAULT)`,
	}

	return commands
}

func GetReais() []string {
	commands := []string{
		`DELETE FROM kudos WHERE id != 0`,
		`DELETE FROM person WHERE id != 0`,
		`ALTER TABLE kudos AUTO_INCREMENT = 1`,
		`ALTER TABLE person AUTO_INCREMENT = 1`,
		`INSERT INTO person (id, name, points, reais) 
			VALUES (DEFAULT, 'José da Silva Santos', 135, 40)`,
		`INSERT INTO kudos (id, personIdFK, name) 
			VALUES (DEFAULT, 1, 'SUPER')`,
		`INSERT INTO kudos (id, personIdFK, name) 
			VALUES (DEFAULT, 1, 'GOOD')`,
		`INSERT INTO kudos (id, personIdFK, name) 
			VALUES (DEFAULT, 1, 'NICE')`,
		`INSERT INTO kudos (id, personIdFK, name) 
			VALUES (DEFAULT, 1, 'OK')`,
	}

	return commands
}

func GetPointsToMessage() []string {
	commands := []string{
		`DELETE FROM kudos WHERE id != 0`,
		`DELETE FROM person WHERE id != 0`,
		`ALTER TABLE kudos AUTO_INCREMENT = 1`,
		`ALTER TABLE person AUTO_INCREMENT = 1`,
		`INSERT INTO person (id, name, points, reais) 
			VALUES (DEFAULT, 'José da Silva Santos', 135, 40)`,
	}

	return commands
}
