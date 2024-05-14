package user

import "context"

type User struct {
	ID int64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserReq struct {
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserRes struct {
	ID int64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
}

type LoginUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	accessToken string
	ID int64 `json:"id"`
	Username string `json:"username"`
}

type UserStore interface {
	createUser(ctx context.Context, user *User) (*User, error)
	getUserByEmail(ctx context.Context, email string) (*User, error)
}