package db

import (
	"database/sql"
)

type Model interface {
	DB() *sql.DB
	GetId() string
	SetId(id string)
	HasId() bool
}
