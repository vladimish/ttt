package domain

import (
	"fmt"
)

func (d *Domain) End(name string, end int64) error {
	// Check if record exists
	r, err := d.j.GetRecord(name)
	if err != nil {
		return fmt.Errorf("failed to get record: %w", err)
	}

	if r.End != 0 {
		return fmt.Errorf("record already ended")
	}

	// Update record
	r.End = end
	err = d.j.SaveRecord(r)
	if err != nil {
		return fmt.Errorf("failed to add record: %w", err)
	}

	return nil
}
