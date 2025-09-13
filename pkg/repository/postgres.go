package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	userTable      = "users"
	todoListsTable = "todo_lists"
	userListsTable = "users_lists"
	todoItemsTable = "todo_items"
	listItemsTable = "list_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode, cfg.Password))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// import (
// 	"database/sql"
// )

// type PostgresDB struct {
// 	db *sql.DB
// }

// func NewPostgresDB(db *sql.DB) *PostgresDB {
// 	return &PostgresDB{db: db}
// }

// func (p *PostgresDB) Close() error {
// 	return p.db.Close()
// }

// func (p *PostgresDB) GetDB() *sql.DB {
// 	return p.db
// }
