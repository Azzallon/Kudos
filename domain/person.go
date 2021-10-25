package domain

type Person struct {
	Id     int    `db:"per_id"`
	Name   string `db:"per_name"`
	Points int    `db:"per_points"`
	Reais  int    `db:"per_reais"`
}
