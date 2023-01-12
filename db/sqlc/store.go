package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store interface {
	Querier
	OrderTx(ctx context.Context, arg OrderTxParams) (OrderTxResult, error)
}

//Store 用于提供数据库操作的各种函数
type SQLStore struct {
	*Queries
	db *sql.DB
}

//用于生成Store结构体
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type OrderTxParams struct {
	UserName    string  `json:"userName"`
	OrderTypeID int64   `json:"orderTypeID"`
	Discount    float64 `json:"discount"`
}

type OrderTxResult struct {
	Order Order `json:"order"`
	User  User  `json:"user"`
}

func (store *SQLStore) OrderTx(ctx context.Context, arg OrderTxParams) (OrderTxResult, error) {
	var result OrderTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		user, err := q.GetUser(ctx, arg.UserName)
		if err != nil {
			return err
		}

		orderType, err := q.GetOrderType(ctx, arg.OrderTypeID)
		if err != nil {
			return err
		}

		//若过期时间在现在时间之前，则从现在起算增加过期时间，否则在过期时间上叠加
		var newExpireTime time.Time
		if user.ExpireAt.Before(time.Now()) {
			newExpireTime = time.Now().AddDate(0, 0, int(orderType.Days))
		} else {
			newExpireTime = user.ExpireAt.AddDate(0, 0, int(orderType.Days))
		}

		result.Order, err = q.CreateOrder(ctx, CreateOrderParams{
			UserName:    arg.UserName,
			OrderTypeID: arg.OrderTypeID,
			Discount:    arg.Discount,
		})
		if err != nil {
			return err
		}

		result.User, err = q.UpdateUserExpireTime(ctx, UpdateUserExpireTimeParams{
			Username: arg.UserName,
			ExpireAt: newExpireTime,
		})

		return nil
	})

	return result, err
}
