package domain

import "fmt"

func (d *Domain) Delete(name string) error {
	err := d.j.DeleteRecord(name)
	if err != nil {
		return fmt.Errorf("can't delete record: %w", err)
	}

	return nil
}
