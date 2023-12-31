package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                 int64  `json:"id" db:"id"`
	Name               string `json:"name" db:"name" validate:"required,min=1,max=30"`
	Surname            string `json:"surname" db:"surname" validate:"required,max=30"`
	Password           string `json:"password" db:"password" validate:"required,min=4"`
	Email              string `json:"email" db:"email" validate:"required,email"`
	Balance            int64  `json:"balance" db:"balance"`
	IsRegisteredSocial bool   `json:"is_register_social" db:"is_register_social"`
	IsPaid             bool   `json:"is_paid" db:"is_paid"`
	Role               string `json:"role" db:"role"` // admin, copywriter, customer
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type UserLogin struct {
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required,min=4"`
}

type UsersList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	Users      []*User `json:"users"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}
