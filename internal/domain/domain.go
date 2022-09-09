package domain

import (
	"github.com/vladimish/ttt/internal/repositories"
)

type Domain struct {
	j repositories.Journal
}

func NewDomain(j repositories.Journal) *Domain {
	return &Domain{
		j: j,
	}
}
