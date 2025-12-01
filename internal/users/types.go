package users

import (
	"context"
	"time"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID          int64      `json:"id"`
	Username    string     `json:"username"`
	Password    string     `json:"-"`
	IsActive    bool       `json:"isActive"`
	Role        Role       `json:"role"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	LastLoginAt *time.Time `json:"lastLoginAt,omitempty"`
}

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type UpdateUserParams struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
	IsActive bool   `json:"isActive"`
}

type Repository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindByID(ctx context.Context, id int64) (User, error)
	Create(ctx context.Context, params CreateUserParams) (User, error)
	Update(ctx context.Context, params UpdateUserParams) (User, error)
}
