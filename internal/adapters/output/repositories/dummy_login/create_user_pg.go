package dummylogin

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// PostgresRepo реализация house.Repository с использованием PostgreSQL.
type PostgresRepo struct {
	db *sql.DB
}

var _ Repository = &PostgresRepo{}

func NewPostgresRepository(dataSourceName string) (*PostgresRepo, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresRepo{db: db}, nil
}

func (r *PostgresRepo) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}

// DummyLogin создает пользователя с заданным статусом
func (r *PostgresRepo) DummyLogin(email string, usertype valueobjects.UserType) (entities.User, error) {
	// Начало транзакции
	tx, err := r.db.Begin()
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback() // Откат транзакции в случае ошибки

	var newUser entities.User

	// Исправленный SQL-запрос
	err = tx.QueryRow(`
        INSERT INTO users (role, email) 
        VALUES ($1, $2) 
        RETURNING user_id, email, role`,
		usertype, email).Scan(&newUser.ID, &newUser.Email, &newUser.Type)
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	// Подтверждение транзакции
	if err := tx.Commit(); err != nil {
		return entities.User{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return newUser, nil
}
