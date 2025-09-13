package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Authorisation interface {
}

type TodoList interface {
}

type TodoIdea interface {
}

type Repository struct {
	Authorisation
	TodoList
	TodoIdea
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}

// func NewPostgresDB() (*sql.DB, error) {
// 	host := viper.GetString("database.host")
// 	port := viper.GetString("database.port")
// 	username := viper.GetString("database.username")
// 	password := viper.GetString("database.password")
// 	dbname := viper.GetString("database.dbname")
// 	sslmode := viper.GetString("database.sslmode")

// 	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
// 		host, port, username, password, dbname, sslmode)

// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open database: %w", err)
// 	}

// 	if err := db.Ping(); err != nil {
// 		return nil, fmt.Errorf("failed to ping database: %w", err)
// 	}

// 	return db, nil
// }
