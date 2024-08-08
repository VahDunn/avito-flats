package subscribe

import (
	"avito-flats/internal/domain/entities"
	"context"
	"database/sql"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewSubscribeRepository(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) SubscribeOnHouse(ctx context.Context, houseID entities.HouseID, email string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO subscriptions (house_id, email) VALUES (?, ?)", houseID, email)
	return err
}
func (r *PostgresRepo) GetSubscribers(ctx context.Context, houseID entities.HouseID) ([]string, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT email FROM subscriptions WHERE house_id = $1", houseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscribers []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		subscribers = append(subscribers, email)
	}

	return subscribers, nil
}
