package request

import "core/db/entity"

type User struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`

	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"omitempty,number"`
}

func (u *User) ToEntity() *entity.User {
	if u == nil {
		return nil
	}

	return &entity.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
	}
}
