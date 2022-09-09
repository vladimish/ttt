package domain

import (
	"fmt"

	"github.com/vladimish/ttt/internal/models"
)

func (d *Domain) List(user string) ([]models.Record, error) {
	recs, err := d.j.GetUserRecords(user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user records: %w", err)
	}

	return recs, nil
}
