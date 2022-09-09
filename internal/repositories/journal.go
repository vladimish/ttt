package repositories

import "github.com/vladimish/ttt/internal/models"

type Journal interface {
	// SaveRecord adds a new record to the journal.
	SaveRecord(record models.Record) error
	// GetUserRecords returns all records for the given user.
	GetUserRecords(user string) ([]models.Record, error)
	// GetRecord returns a record by its name.
	GetRecord(name string) (models.Record, error)
}
