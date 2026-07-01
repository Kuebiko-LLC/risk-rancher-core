package datastore

import (
	"database/sql"

	"code.riskrancher.com/RiskRancher/core/pkg/domain"
)

type SQLiteStore struct {
	DB *sql.DB
}

var _ domain.TicketStore = (*SQLiteStore)(nil)

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{DB: db}
}
