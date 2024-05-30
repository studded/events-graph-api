package model

type Event struct {
	ID          int `goqu:"skipinsert"`
	Name        string
	StartDate   string `db:"start_date"`
	EndDate     string `db:"end_date"`
	Location    string
	Description string
}

type NewEvent struct {
	UserID      int `db:"user_id"`
	Name        string
	StartDate   string `db:"start_date"`
	EndDate     string `db:"end_date"`
	Location    string
	Description string
}
