package model

type Expense struct {
	ID          int `goqu:"skipinsert,skipupdate"`
	Event       *Event
	EventID     int `db:"event_id"`
	Name        string
	Cost        float64
	Description string
	Category    string
}

type NewExpense struct {
	EventID     int
	Name        string
	Cost        float64
	Description string
	Category    string
}

type UpdateExpense struct {
	Name        *string
	Cost        *float64
	Description *string
	Category    *string
}

type ExpenseCategory struct {
	Category string
	Total    float64
}
