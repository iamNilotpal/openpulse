package organizations_store

import (
	"github.com/jmoiron/sqlx"
)

type Store interface {
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}
