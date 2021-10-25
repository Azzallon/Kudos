package domain

type Kudo struct {
	Id         int    `db:"kudo_id"`
	PersonIdFK int    `db:"kudo_per_id_fk"`
	Name       string `db:"kudo_name"`
}
