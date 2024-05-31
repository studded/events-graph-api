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

func (e *EventsRepo) GetEvents(limit, offset *int) ([]*model.Event, error) {
	// SELECT * FROM events LIMIT {limit} OFFSET {offset}
	query, _, _ := goqu.From("events").Limit(uint(*limit)).Offset(uint(*offset)).ToSQL()
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

func (e *EventsRepo) UpdateEvent(event *model.Event) (*model.Event, error) {
	// UPDATE events SET {col}={val}, ... WHERE id={id} RETURNING *
	query, _, _ := goqu.Update("events").Set(event).Where(goqu.C("id").Eq(event.ID)).Returning("*").ToSQL()
	fmt.Println(query)

	err := e.DB.Get(event, query)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (e *EventsRepo) DeleteEvent(event *model.Event) error {
	tx := e.DB.MustBegin()

	// DELETE FROM events WHERE id={id}
	query, _, _ := goqu.Delete("events").Where(goqu.C("id").Eq(event.ID)).ToSQL()
	fmt.Println(query)
	tx.MustExec(query)

	// DELETE FROM activities WHERE event_id={event.id}
	query, _, _ = goqu.Delete("activities").Where(goqu.C("event_id").Eq(event.ID)).ToSQL()
	fmt.Println(query)
	tx.MustExec(query)

	// DELETE FROM roles WHERE event_id={event.id}
	query, _, _ = goqu.Delete("roles").Where(goqu.C("event_id").Eq(event.ID)).ToSQL()
	fmt.Println(query)
	tx.MustExec(query)

	// DELETE FROM expenses WHERE event_id={event.id}
	query, _, _ = goqu.Delete("expenses").Where(goqu.C("event_id").Eq(event.ID)).ToSQL()
	fmt.Println(query)
	tx.MustExec(query)

	return tx.Commit()
}
