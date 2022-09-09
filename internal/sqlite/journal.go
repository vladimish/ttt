package sqlite

import (
	"fmt"

	"github.com/vladimish/ttt/internal/models"
)

var req = `INSERT OR REPLACE INTO records (id, USER, NAME, description, START, END)	VALUES ((SELECT ID FROM records WHERE NAME = :name), :user, :name, :description, :start, :end);`

// SaveRecord adds a new record to the journal.
func (s *Sqlite) SaveRecord(record models.Record) error {
	_, err := s.conn.NamedExec(
		fmt.Sprintf(req),
		record,
	)
	if err != nil {
		return fmt.Errorf("failed to insert record: %w", err)
	}

	return nil
}

// GetUserRecords returns all records for the given user.
func (s *Sqlite) GetUserRecords(user string) (res []models.Record, err error) {
	resp, err := s.conn.Query(`SELECT id, USER, name, description, start, "end" FROM records WHERE USER=?;`, user)
	if err != nil {
		return nil, fmt.Errorf("failed to select records: %w", err)
	}

	for resp.Next() {
		t := models.Record{}
		err = resp.Scan(&t.ID, &t.User, &t.Name, &t.Description, &t.Start, &t.End)
		if err != nil {
			return nil, fmt.Errorf("failed to scan records: %w", err)
		}

		res = append(res, t)
	}
	if resp.Err() != nil {
		return nil, fmt.Errorf("failed to iterate over response: %w", err)
	}

	return res, nil
}

// GetRecord returns a record by its name.
func (s *Sqlite) GetRecord(name string) (res models.Record, err error) {
	err = s.conn.Get(&res, `SELECT id, USER, name, description, start, "end" FROM records WHERE name=?;`, name)
	if err != nil {
		return models.Record{}, fmt.Errorf("failed to select record: %w", err)
	}

	return res, nil
}
