package user

import (
	"context"
	"core/db/entity"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository - repository interface for working with users
type Repository interface {
	Create(ctx context.Context, u *entity.User) error
	GetByID(ctx context.Context, id int) (*entity.User, error)
	List(ctx context.Context) ([]*entity.User, error)
	UpdateOne(ctx context.Context, u *entity.User) error
	UpdateMany(ctx context.Context, u *entity.User) error
	Delete(ctx context.Context, id int) error
}

type DBRepo struct {
	db *pgxpool.Pool
}

func NewDBRepo(db *pgxpool.Pool) *DBRepo {
	return &DBRepo{db: db}
}

// Create добавляет нового пользователя в базу данных
func (r *DBRepo) Create(ctx context.Context, u *entity.User) error {
	const query = `INSERT INTO users (first_name, last_name, email, phone) VALUES ($1, $2, $3, $4) RETURNING id`

	var userID int
	if err := r.db.QueryRow(ctx, query, u.FirstName, u.LastName, u.Email, u.Phone).Scan(&userID); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Присвоение ID вновь созданного пользователя
	u.ID = userID

	return nil
}

func (r *DBRepo) GetByID(ctx context.Context, id int) (*entity.User, error) {
	return nil, nil
}

// List возвращает список пользователей из базы данных
func (r *DBRepo) List(ctx context.Context) ([]*entity.User, error) {
	const query = `SELECT id, first_name, last_name, email, phone FROM users`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	defer rows.Close()

	var users []*entity.User

	// Перебираем строки результата и заполняем слайс пользователей
	for rows.Next() {
		var u entity.User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Phone); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &u)
	}

	// Проверка на ошибки после перебора
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during rows iteration: %w", err)
	}

	return users, nil
}

func (r *DBRepo) UpdateOne(ctx context.Context, u *entity.User) error {
	return nil
}

func (r *DBRepo) UpdateMany(ctx context.Context, u *entity.User) error {
	return nil
}

func (r *DBRepo) Delete(ctx context.Context, id int) error {
	return nil
}
