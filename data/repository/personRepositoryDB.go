package repository

import (
	"database/sql"
	"kudos/data/db"
	"kudos/domain"
	"kudos/domain/errs"
	"kudos/utils"
)

type PersonRepositoryDB struct{}

func NewPersonRepositoryDB() *PersonRepositoryDB {
	return &PersonRepositoryDB{}
}

func (p PersonRepositoryDB) GetPeople() ([]domain.Person, *errs.ApiError) {
	db := db.GetConnection()

	selDB, err := db.Query("SELECT * FROM person")
	if err != nil {
		return nil, errs.GetGenericErrorMessage(err)
	}

	person := domain.Person{}
	personList := []domain.Person{}

	for selDB.Next() {
		var (
			id     int
			name   string
			points int
			reais  int
		)

		err = selDB.Scan(&id, &name, &points, &reais)
		if err != nil {
			return nil, errs.GetGenericErrorMessage(err)
		}

		person.Id = id
		person.Name = name
		person.Points = points
		person.Reais = reais

		personList = append(personList, person)
	}

	defer db.Close()

	return personList, nil
}

func (p PersonRepositoryDB) CreatePerson(name string) *errs.ApiError {
	db := db.GetConnection()

	stmt, err := db.Prepare("INSERT INTO person (name) VALUES (?)")
	if err != nil {
		return errs.GetGenericErrorMessage(err)
	}

	stmt.Exec(name)

	defer db.Close()

	return nil
}

func (p PersonRepositoryDB) AddPoints(id, points int) *errs.ApiError {
	db := db.GetConnection()

	result, err := db.Exec(
		"UPDATE person SET points = points + ? WHERE id = ?", points, id)
	if err != nil {
		return errs.GetGenericErrorMessage(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errs.GetGenericErrorMessage(err)
	}

	if rowsAffected == 0 {
		return errs.NewBadRequestError("A pessoa informada não existe!")
	} else {

		apiErr := insertKudos(db, points, id)
		if apiErr != nil {
			return apiErr
		}

		apiErr = updateReais(db, points, id)
		if apiErr != nil {
			return apiErr
		}

		return nil
	}
}

func insertKudos(db *sql.DB, points, id int) *errs.ApiError {
	kudosList := utils.GetKudosList(points)

	for i := 0; i < len(kudosList); i++ {
		_, err := db.Exec(
			"INSERT INTO kudos (personIdFK, name) VALUES (?, ?)", id, kudosList[i])
		if err != nil {
			return errs.GetGenericErrorMessage(err)
		}
	}

	return nil
}

func updateReais(db *sql.DB, points, id int) *errs.ApiError {
	reaisSum := utils.GetReaisSum(points)

	_, err := db.Exec(
		"UPDATE person SET reais = reais + ? WHERE id = ?", reaisSum, id)
	if err != nil {
		return errs.GetGenericErrorMessage(err)
	}

	return nil
}

func (p PersonRepositoryDB) GetReais(id int) (int, *errs.ApiError) {
	db := db.GetConnection()

	selDB, err := db.Query("SELECT reais FROM person WHERE id = ?", id)
	if err != nil {
		return 0, errs.GetGenericErrorMessage(err)
	}

	personReais := -1

	for selDB.Next() {
		var reais int

		err = selDB.Scan(&reais)
		if err != nil {
			return 0, errs.GetGenericErrorMessage(err)
		}

		personReais = reais
	}

	if personReais == -1 {
		return 0, errs.NewBadRequestError("A pessoa informada não existe!")
	}

	defer db.Close()

	return personReais, nil
}

func (p PersonRepositoryDB) GetPointsToMessage(id int) (string, *errs.ApiError) {
	db := db.GetConnection()

	selDB, err := db.Query("SELECT points FROM person WHERE id = ?", id)
	if err != nil {
		return "", errs.GetGenericErrorMessage(err)
	}

	personPoints := -1

	for selDB.Next() {
		var points int

		err = selDB.Scan(&points)
		if err != nil {
			return "", errs.GetGenericErrorMessage(err)
		}

		personPoints = points
	}

	if personPoints == -1 {
		return "", errs.NewBadRequestError("A pessoa informada não existe!")
	}

	msg := utils.GetConvertPointsToMessage(personPoints)

	defer db.Close()

	return msg, nil
}
