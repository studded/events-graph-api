package postgres

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"fmt"

	"github.com/studded/events-graph-api/graph/model"
)

type ExpensesRepo struct {
	DB *sqlx.DB
}

func (e *ExpensesRepo) GetExpenseByID(id int) (*model.Expense, error) {
	// SELECT * FROM expenses WHERE id={id}
	query, _, _ := goqu.From("expenses").Where(goqu.C("id").Eq(id)).ToSQL()
	fmt.Println(query)

	var expense model.Expense
	err := e.DB.Get(&expense, query)
	if err != nil {
		return nil, err
	}

	return &expense, nil
}

func (e *ExpensesRepo) GetExpensesByEventID(eventId int) ([]*model.Expense, error) {
	// SELECT * FROM expenses WHERE event_id={eventId}
	query, _, _ := goqu.From("expenses").Where(goqu.C("event_id").Eq(eventId)).ToSQL()
	fmt.Println(query)

	var expenses []*model.Expense
	err := e.DB.Select(&expenses, query)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}
