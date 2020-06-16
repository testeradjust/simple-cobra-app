package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DummyPostgres struct {
	db *sql.DB
}

func NewDummyPostgres(host, user, dbName string, port int) (*DummyPostgres, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres conn: %s", err)
	}

	return &DummyPostgres{
		db: db,
	}, nil
}

func (dp *DummyPostgres) TestConnection() error {
	return dp.db.Ping()
}
