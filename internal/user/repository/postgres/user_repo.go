package postgres

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/user"
	"github.com/ivanwang123/roadmap/models"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) user.Repository {
	return &userRepo{db}
}

func (r *userRepo) GetAll(ctx context.Context) ([]*models.User, error) {
	users := []*models.User{}
	if err := r.db.Select(&users, "SELECT * FROM users"); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) GetByID(ctx context.Context, ID int) (*models.User, error) {
	var user models.User
	if err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1 LIMIT 1", ID); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1 LIMIT 1", email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := r.db.Get(&user, "SELECT * FROM users WHERE username = $1 LIMIT 1", username); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Create(ctx context.Context, input *models.NewUser) (*models.User, error) {
	var user models.User
	if err := r.db.Get(&user, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *",
		input.Username, input.Email, input.Password); err != nil {
		return nil, err
	}
	return &user, nil
}
