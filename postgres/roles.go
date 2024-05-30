package postgres

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"fmt"

	"github.com/studded/events-graph-api/graph/model"
)

type RolesRepo struct {
	DB *sqlx.DB
}

func (e *RolesRepo) GetRoleByID(id int) (*model.Role, error) {
	// SELECT * FROM roles WHERE id={id}
	query, _, _ := goqu.From("roles").Where(goqu.C("id").Eq(id)).ToSQL()
	fmt.Println(query)

	var role model.Role
	err := e.DB.Get(&role, query)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (e *RolesRepo) GetRolesByEventID(eventId int) ([]*model.Role, error) {
	// SELECT * FROM roles WHERE event_id={eventId}
	query, _, _ := goqu.From("roles").Where(goqu.C("event_id").Eq(eventId)).ToSQL()
	fmt.Println(query)

	var roles []*model.Role
	err := e.DB.Select(&roles, query)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (e *RolesRepo) GetRolesByUserID(userId int) ([]*model.Role, error) {
	// SELECT * FROM roles WHERE event_id={eventId}
	query, _, _ := goqu.From("roles").Where(goqu.C("user_id").Eq(userId)).ToSQL()
	fmt.Println(query)

	var roles []*model.Role
	err := e.DB.Select(&roles, query)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (e *RolesRepo) GetRoleByEventIDAndUserID(eventId int, userId int) (*model.Role, error) {
	// SELECT * FROM roles WHERE user_id={userId} AND event_id={eventId}
	query, _, _ := goqu.From("roles").Where(
		goqu.C("user_id").Eq(userId),
		goqu.C("event_id").Eq(eventId),
	).ToSQL()
	fmt.Println(query)

	var role model.Role
	err := e.DB.Get(&role, query)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (e *RolesRepo) CreateRole(role *model.Role) (*model.Role, error) {
	// INSERT INTO roles {cols} VALUES {vals} RETURNING *
	query, _, _ := goqu.Insert("roles").Rows(role).Returning("*").ToSQL()
	fmt.Println(query)

	err := e.DB.Get(role, query)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (e *RolesRepo) UpdateRole(role *model.Role) (*model.Role, error) {
	// UPDATE roles SET {col}={val}, ... WHERE id={id} RETURNING *
	query, _, _ := goqu.Update("roles").Set(role).Where(goqu.C("id").Eq(role.ID)).Returning("*").ToSQL()
	fmt.Println(query)

	err := e.DB.Get(role, query)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (e *RolesRepo) DeleteRole(role *model.Role) error {
	// DELETE FROM roles WHERE id={id}
	query, _, _ := goqu.Delete("roles").Where(goqu.C("id").Eq(role.ID)).ToSQL()
	fmt.Println(query)

	_, err := e.DB.Exec(query)

	return err
}
