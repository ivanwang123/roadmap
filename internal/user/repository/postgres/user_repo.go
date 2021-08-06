package postgres

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/common/transaction"
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
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Select(&users, "SELECT * FROM users"); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) GetByID(ctx context.Context, ID int) (*models.User, error) {
	var user models.User
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Get(&user, "SELECT * FROM users WHERE id = $1 LIMIT 1", ID); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Get(&user, "SELECT * FROM users WHERE email = $1 LIMIT 1", email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Get(&user, "SELECT * FROM users WHERE username = $1 LIMIT 1", username); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetByRoadmapFollowing(ctx context.Context, roadmapID int) ([]*models.User, error) {
	users := []*models.User{}
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Select(&users, "SELECT u.* FROM roadmap_followers AS f LEFT JOIN users AS u ON f.user_id = u.id WHERE f.roadmap_id = $1", roadmapID); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) GetIn(ctx context.Context, IDs []string) ([]*models.User, error) {
	users := []*models.User{}
	query, args, err := sqlx.In("SELECT * FROM users WHERE id IN (?)", IDs)
	if err != nil {
		return nil, err
	}

	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Select(&users, r.db.Rebind(query), args...); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) Create(ctx context.Context, input *models.NewUser) (*models.User, error) {
	var user models.User
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Get(&user, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *",
		input.Username, input.Email, input.Password); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) WithTransaction(ctx context.Context, fn transaction.TxFunc) error {
	return transaction.NewTransaction(ctx, r.db, fn)
}
