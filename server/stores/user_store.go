package stores

import (
	"fmt"

	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	*sqlx.DB
}

func (s *UserStore) Create(input *model.NewUser) (*model.User, error) {
	var user model.User
	if err := s.Get(&user, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *",
		input.Username, input.Email, input.Password); err != nil {
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

// TODO: Deprecated?
func (s *UserStore) GetById(id int) (*model.User, error) {
	var user model.User
	if err := s.Get(&user, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}
	fmt.Println("GET USER BY ID")
	return &user, nil
}
