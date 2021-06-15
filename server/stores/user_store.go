package stores

import (
	"errors"
	"fmt"

	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	*sqlx.DB
}

func (s *UserStore) Create(input *model.NewUser) (*model.User, error) {
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	var user model.User
	if err := s.Get(&user, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *",
		input.Username, input.Email, hashedPassword); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) GetAll() ([]*model.User, error) {
	users := []*model.User{}
	if err := s.Select(&users, "SELECT * FROM users"); err != nil {
		return []*model.User{}, err
	}
	fmt.Println("GET ALL USERS")
	return users, nil
}

func (s *UserStore) GetById(id int) (*model.User, error) {
	var user model.User
	if err := s.Get(&user, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) Authenticate(input *model.Login) (int, error) {
	var user model.User
	if input.Email != nil {
		if err := s.Get(&user, "SELECT * FROM users WHERE email = $1", input.Email); err != nil {
			return -1, errors.New("Invalid email")
		}
	} else if input.Username != nil {
		if err := s.Get(&user, "SELECT * FROM users WHERE username = $1", input.Username); err != nil {
			return -1, errors.New("Invalid username")
		}
	}

	if CheckHashedPassword(input.Password, user.Password) {
		return user.ID, nil
	} else {
		return -1, errors.New("Invalid password")
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHashedPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
