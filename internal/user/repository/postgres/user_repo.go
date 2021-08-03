package postgres

import (
	"context"
	"errors"

	"github.com/ivanwang123/roadmap/internal/user"
	"github.com/ivanwang123/roadmap/internal/user/repository"
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

func (r *userRepo) Create(ctx context.Context, input *models.NewUser) (*models.User, error) {
	hashedPassword, err := repository.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := r.db.Get(&user, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *",
		input.Username, input.Email, hashedPassword); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Authenticate(ctx context.Context, input *models.Login) (*models.User, error) {
	var user models.User
	var err error
	if input.Email != nil {
		err = r.db.Get(&user, "SELECT * FROM users WHERE email = $1 LIMIT 1", input.Email)
	} else if input.Username != nil {
		err = r.db.Get(&user, "SELECT * FROM users WHERE username = $1 LIMIT 1", input.Username)
	} else {
		err = errors.New("Missing credentials")
	}

	if err != nil {
		return nil, err
	}

	if !repository.CheckPasswordHash(user.Password, input.Password) {
		return nil, errors.New("Incorrect password")
	}

	return &user, nil
}
