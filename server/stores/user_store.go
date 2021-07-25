package stores

import (
	"errors"

	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	DB *sqlx.DB
}

func (s *UserStore) Create(input *model.NewUser) (*model.User, error) {
	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	var user model.User
	if err := s.DB.Get(&user, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *",
		input.Username, input.Email, hashedPassword); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) GetAll() ([]*model.User, error) {
	users := []*model.User{}
	if err := s.DB.Select(&users, "SELECT * FROM users"); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserStore) GetById(id int) (*model.User, error) {
	var user model.User
	if err := s.DB.Get(&user, "SELECT * FROM users WHERE id = $1 LIMIT 1", id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) Authenticate(input *model.Login) (*model.User, error) {
	var user model.User
	var err error
	if input.Email != nil {
		err = s.DB.Get(&user, "SELECT * FROM users WHERE email = $1 LIMIT 1", input.Email)
	} else if input.Username != nil {
		err = s.DB.Get(&user, "SELECT * FROM users WHERE username = $1 LIMIT 1", input.Username)
	} else {
		err = errors.New("Missing credentials")
	}

	if err != nil {
		return nil, err
	}

	if !checkPasswordHash(user.Password, input.Password) {
		return nil, errors.New("Incorrect password")
	}

	return &user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
