package repository

import (
	"gitlab.com/kanya384/gotools/psql"
)

type Repository struct {
	*psql.Postgres
}

type Options struct {
}

func New(pg *psql.Postgres, options Options) (*Repository, error) {
	var r = &Repository{pg}
	r.SetOptions(options)
	return r, nil
}

func (r *Repository) SetOptions(options Options) {
}
