package entity

type User struct {
	ID        int    `id:"id" json:"id"`
	FirstName string `id:"first_name" json:"first_name"`
	LastName  string `id:"last_name" json:"last_name"`
	Email     string `id:"email" json:"email"`
	Phone     string `id:"phone" json:"phone"`
}
