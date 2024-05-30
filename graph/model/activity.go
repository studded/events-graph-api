package model

type Activity struct {
	ID          int `goqu:"skipinsert,skipupdate"`
	Event       *Event
	EventID     int `db:"event_id"`
	Name        string
	StartTime   string `db:"start_time"`
	EndTime     string `db:"end_time"`
	Description string
}

type NewActivity struct {
	EventID     int
	Name        string
	StartTime   string
	EndTime     string
	Description string
}

type UpdateActivity struct {
	Name        *string
	StartTime   *string
	EndTime     *string
	Description *string
}
