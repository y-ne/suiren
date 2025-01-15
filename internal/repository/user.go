package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"suiren/internal/model"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user model.CreateUserRequest) error {
	_, err := r.db.Exec(ctx,
		"INSERT INTO users (email, username, password) VALUES ($1, $2, $3)",
		user.Email, user.Username, user.Password)
	return err
}

func (r *UserRepo) List(ctx context.Context) ([]model.User, error) {
	rows, err := r.db.Query(ctx, "SELECT id, email, username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Email, &u.Username); err != nil {
			continue
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepo) GetByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow(ctx,
		"SELECT id, email, username FROM users WHERE id = $1",
		id).Scan(&user.ID, &user.Email, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
