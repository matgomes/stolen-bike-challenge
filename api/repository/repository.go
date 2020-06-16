package repository

import (
    "database/sql"
    "fmt"
    "github.com/matgomes/stolen-bike-challenge/config"
)

type Repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{
        db: db,
    }
}

func Connect(c config.DBConfig) (*sql.DB, error) {
    dsn := fmt.Sprintf("host=%s port=5432 dbname=%s user=postgres password='root' sslmode=disable", c.Server, c.Name)
    return sql.Open("postgres", dsn)
}

func (r *Repository) CloseConn() error {
    return r.db.Close()
}

type scanner interface {
    Scan(...interface{}) error
}

func getNullableID(id int) sql.NullInt64 {
    return sql.NullInt64{
        Int64: int64(id),
        Valid: id > 0,
    }
}
