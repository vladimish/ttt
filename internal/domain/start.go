package domain

import (
	"fmt"

	"github.com/vladimish/ttt/internal/models"
)

func (d *Domain) Start(name string, description string, start int64) error {
	record := models.Record{
		Name:        name,
		Start:       start,
		Description: description,
	}

	err := d.j.SaveRecord(record)
	if err != nil {
		return fmt.Errorf("failed to add record: %w", err)
	}

	return nil
}
