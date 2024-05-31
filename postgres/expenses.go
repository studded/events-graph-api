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

func (e *ExpensesRepo) CreateExpense(expense *model.Expense) (*model.Expense, error) {
	// INSERT INTO expenses {cols} VALUES {vals} RETURNING *
	query, _, _ := goqu.Insert("expenses").Rows(expense).Returning("*").ToSQL()
	fmt.Println(query)

	err := e.DB.Get(expense, query)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (e *ExpensesRepo) UpdateExpense(expense *model.Expense) (*model.Expense, error) {
	// UPDATE expenses SET {col}={val}, ... WHERE id={id} RETURNING *
	query, _, _ := goqu.Update("expenses").Set(expense).Where(goqu.C("id").Eq(expense.ID)).Returning("*").ToSQL()
	fmt.Println(query)

	err := e.DB.Get(expense, query)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (e *ExpensesRepo) DeleteExpense(expense *model.Expense) error {
	// DELETE FROM expenses WHERE id={id}
	query, _, _ := goqu.Delete("expenses").Where(goqu.C("id").Eq(expense.ID)).ToSQL()
	fmt.Println(query)

	_, err := e.DB.Exec(query)

	return err
}

func (e *ExpensesRepo) GetExpensesTotalByEventID(eventId int) (float64, error) {
	// SELECT COALESCE(SUM(cost), 0) FROM expenses WHERE event_id={eventId}
	query, _, _ := goqu.Select(goqu.COALESCE(goqu.SUM("cost"), 0)).
		From("expenses").Where(goqu.C("event_id").Eq(eventId)).ToSQL()
	fmt.Println(query)

	var total float64
	err := e.DB.Get(&total, query)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (e *ExpensesRepo) GetExpenseCategoriesByEventID(eventId int) ([]*model.ExpenseCategory, error) {
	// SELECT category, SUM(cost) as total FROM expenses WHERE event_id={event_id} GROUP BY category
	query, _, _ := goqu.Select("category", goqu.COUNT("*"), goqu.SUM("cost").As("total")).
		From("expenses").Where(goqu.C("event_id").Eq(eventId)).GroupBy("category").ToSQL()
	fmt.Println(query)

	var expenseCategories []*model.ExpenseCategory
	err := e.DB.Select(&expenseCategories, query)
	if err != nil {
		return nil, err
	}

	return expenseCategories, nil
}
