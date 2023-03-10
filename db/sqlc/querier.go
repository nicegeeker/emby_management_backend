// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	CreateOrderType(ctx context.Context, arg CreateOrderTypeParams) (OrderType, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, username string) error
	GetOrder(ctx context.Context, id int64) (Order, error)
	GetOrderType(ctx context.Context, id int64) (OrderType, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListOrderTypes(ctx context.Context) ([]OrderType, error)
	ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateUserExpireTime(ctx context.Context, arg UpdateUserExpireTimeParams) (User, error)
}

var _ Querier = (*Queries)(nil)
