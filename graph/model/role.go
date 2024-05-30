package model

type Role struct {
	ID      int `goqu:"skipinsert"`
	User    *User
	UserID  int `db:"user_id"`
	Event   *Event
	EventID int `db:"event_id"`
	Type    string
}

type NewRole struct {
	UserID  int
	EventID int
	Type    string
}
