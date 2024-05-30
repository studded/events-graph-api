package postgres

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"fmt"

	"github.com/studded/events-graph-api/graph/model"
)

type ActivitiesRepo struct {
	DB *sqlx.DB
}

func (e *ActivitiesRepo) GetActivityByID(id int) (*model.Activity, error) {
	// SELECT * FROM activities WHERE id={id}
	query, _, _ := goqu.From("activities").Where(goqu.C("id").Eq(id)).ToSQL()
	fmt.Println(query)

	var activity model.Activity
	err := e.DB.Get(&activity, query)
	if err != nil {
		return nil, err
	}

	return &activity, nil
}

func (e *ActivitiesRepo) GetActivitiesByEventID(eventId int) ([]*model.Activity, error) {
	// SELECT * FROM activities WHERE event_id={eventId}
	query, _, _ := goqu.From("activities").Where(goqu.C("event_id").Eq(eventId)).ToSQL()
	fmt.Println(query)

	var activities []*model.Activity
	err := e.DB.Select(&activities, query)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (e *ActivitiesRepo) CreateActivity(activity *model.Activity) (*model.Activity, error) {
	// INSERT INTO activities {cols} VALUES {vals} RETURNING *
	query, _, _ := goqu.Insert("activities").Rows(activity).Returning("*").ToSQL()
	fmt.Println(query)

	err := e.DB.Get(activity, query)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (e *ActivitiesRepo) UpdateActivity(activity *model.Activity) (*model.Activity, error) {
	// UPDATE activities SET {col}={val}, ... WHERE id={id} RETURNING *
	query, _, _ := goqu.Update("activities").Set(activity).Where(goqu.C("id").Eq(activity.ID)).Returning("*").ToSQL()
	fmt.Println(query)

	err := e.DB.Get(activity, query)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (e *ActivitiesRepo) DeleteActivity(activity *model.Activity) error {
	// DELETE FROM activities WHERE id={id}
	query, _, _ := goqu.Delete("activities").Where(goqu.C("id").Eq(activity.ID)).ToSQL()
	fmt.Println(query)

	_, err := e.DB.Exec(query)

	return err
}
