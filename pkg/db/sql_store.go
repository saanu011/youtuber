package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	// register postgres driver
	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
)

const (
	ErrInvalidTransaction   = "no valid transaction"
	ErrCantStartTransaction = "can't start transaction"
	ErrCantCloseTransaction = "can't close transaction"
)

const (
	newrelicPostgresDriver = "nrpostgres"
)

type SqlDB struct {
	*sqlx.DB
}

type SqlTx struct {
	*sqlx.Tx
}

func NewSQLStore(dbConfig Config) (*SqlDB, error) {
	db, err := sqlx.Connect(newrelicPostgresDriver, dbConfig.ConnectionString())
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(dbConfig.MaxIdleConnections)
	db.SetMaxOpenConns(dbConfig.MaxOpenConnections)
	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetimeMin) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(dbConfig.ConnMaxIdleTime) * time.Millisecond)
	sqlDB := &SqlDB{db}

	return sqlDB, nil
}
