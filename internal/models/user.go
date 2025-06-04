package models

type User struct {
	ID        int64
	Login     string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type GetUserFilter struct {
	Id        *int64
	Login     *string
	FirstName *string
	LastName  *string
	Email     *string
	Phone     *string
}
