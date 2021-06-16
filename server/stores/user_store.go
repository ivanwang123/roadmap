package stores

import (
	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	*sqlx.DB
}

func (s *UserStore) Create(input *model.NewUser) (*model.User, error) {
	var user model.User
	if err := s.Get(&user, "INSERT INTO users (username, email) VALUES ($1, $2) RETURNING *",
		input.Username, input.Email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) GetAll() ([]*model.User, error) {
	users := []*model.User{}
	if err := s.Select(&users, "SELECT * FROM users"); err != nil {
		return []*model.User{}, err
	}
	return users, nil
}

func (s *UserStore) GetById(id int) (*model.User, error) {
	var user model.User
	if err := s.Get(&user, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}
	return &user, nil
}
