// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Activity struct {
	ID          string `json:"id"`
	EventID     int    `json:"event_id"`
	Name        string `json:"name"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Description string `json:"description"`
}

type Event struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	StartDate   string      `json:"start_date"`
	EndDate     string      `json:"end_date"`
	Location    string      `json:"location"`
	Description string      `json:"description"`
	Activities  []*Activity `json:"activities"`
}

type Mutation struct {
}

type NewActivity struct {
	EventID     string `json:"eventId"`
	Name        string `json:"name"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Description string `json:"description"`
}

type NewUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Query struct {
}

type Roles struct {
	ID          string `json:"id"`
	UserID      int    `json:"user_id"`
	EventID     int    `json:"event_id"`
	Name        string `json:"name"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Description string `json:"description"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}