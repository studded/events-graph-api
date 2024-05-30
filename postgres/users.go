package postgres

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"fmt"

	"github.com/studded/events-graph-api/graph/model"
)

type UsersRepo struct {
	DB *sqlx.DB
}

func (e *UsersRepo) GetUsers() ([]*model.User, error) {
	// SELECT * FROM users
	query, _, _ := goqu.From("users").ToSQL()
	fmt.Println(query)

	var users []*model.User
	err := e.DB.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (e *UsersRepo) GetUserByID(id int) (*model.User, error) {
	// SELECT * FROM users WHERE id={id}
	query, _, _ := goqu.From("users").Where(goqu.C("id").Eq(id)).ToSQL()
	fmt.Println(query)

	var user model.User
	err := e.DB.Get(&user, query)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (e *UsersRepo) GetUserByEmail(email string) (*model.User, error) {
	// SELECT * FROM users WHERE id={id}
	query, _, _ := goqu.From("users").Where(goqu.C("email").Eq(email)).ToSQL()
	fmt.Println(query)

	var user model.User
	err := e.DB.Get(&user, query)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (e *UsersRepo) CreateUser(user *model.User) (*model.User, error) {
	// INSERT INTO users {cols} VALUES {vals} RETURNING *
	query, _, _ := goqu.Insert("users").Rows(user).Returning("*").ToSQL()
	fmt.Println(query)

	err := e.DB.Get(user, query)
	if err != nil {
		return nil, err
	}

	return user, nil
}
