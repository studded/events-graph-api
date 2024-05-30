package postgres

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"fmt"

	"github.com/studded/events-graph-api/graph/model"
)

type EventsRepo struct {
	DB *sqlx.DB
}

func (e *EventsRepo) GetEvents() ([]*model.Event, error) {
	// SELECT * FROM events
	query, _, _ := goqu.From("events").ToSQL()
	fmt.Println(query)

	var events []*model.Event
	err := e.DB.Select(&events, query)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (e *EventsRepo) GetEventByID(id int) (*model.Event, error) {
	// SELECT * FROM events WHERE id={id}
	query, _, _ := goqu.From("events").Where(goqu.C("id").Eq(id)).ToSQL()
	fmt.Println(query)

	var event model.Event
	err := e.DB.Get(&event, query)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *EventsRepo) CreateEvent(event *model.Event) (*model.Event, error) {
	// INSERT INTO events {cols} VALUES {vals} RETURNING *
	query, _, _ := goqu.Insert("events").Rows(event).Returning("*").ToSQL()
	fmt.Println(query)

	err := e.DB.Get(event, query)
	if err != nil {
		return nil, err
	}

	return event, nil
}
