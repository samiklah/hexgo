package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection faliur: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping faliur: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close faliur: %v", err)
	}
}
func (da Adapter) AddToHistory(answer int, operation string) error {
	queryString, args, err := sq.Insert("history").Columns("data", "answer", "operation").
		Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	return nil
}
