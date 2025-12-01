package postgres

import (
	"context"

	repo "github.com/choiexe1/go-ecom/internal/adapters/postgresql/sqlc"
	"github.com/choiexe1/go-ecom/internal/users"
)

type repository struct {
	queries *repo.Queries
}

func NewRepository(queries *repo.Queries) users.Repository {
	return &repository{
		queries: queries,
	}
}

func (r *repository) FindAll(ctx context.Context) ([]users.User, error) {
	rows, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]users.User, len(rows))
	for i, row := range rows {
		result[i] = toUser(row)
	}
	return result, nil
}

func (r *repository) FindByID(ctx context.Context, id int64) (users.User, error) {
	user, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return users.User{}, err
	}

	return toUser(user), nil
}

func (r *repository) Create(ctx context.Context, params users.CreateUserParams) (users.User, error) {
	user, err := r.queries.CreateUser(ctx, repo.CreateUserParams{
		Username: params.Username,
		Password: params.Password,
		Role:     string(params.Role),
	})
	if err != nil {
		return users.User{}, err
	}

	return toUser(user), nil
}

func (r *repository) Update(ctx context.Context, params users.UpdateUserParams) (users.User, error) {
	user, err := r.queries.UpdateUser(ctx, repo.UpdateUserParams{
		ID:       params.ID,
		Password: params.Password,
		Role:     string(params.Role),
		IsActive: params.IsActive,
	})
	if err != nil {
		return users.User{}, err
	}

	return toUser(user), nil
}

func toUser(u repo.User) users.User {
	return users.User{
		ID:          u.ID,
		Username:    u.Username,
		IsActive:    u.IsActive,
		Role:        users.Role(u.Role),
		CreatedAt:   u.CreatedAt.Time,
		UpdatedAt:   u.UpdatedAt.Time,
		LastLoginAt: &u.LastLoginAt.Time,
	}
}
