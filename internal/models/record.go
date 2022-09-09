package models

type Record struct {
	ID          int32  `db:"id"`
	User        string `db:"user"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Start       int64  `db:"start"`
	End         int64  `db:"end"`
}
