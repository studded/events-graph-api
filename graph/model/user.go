package model

type User struct {
	ID    int `goqu:"skipinsert"`
	Name  string
	Email string
	Phone string
}

type NewUser struct {
	Name  string
	Email string
	Phone string
}
