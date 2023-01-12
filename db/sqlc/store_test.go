package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrderTx(t *testing.T) {
	store := NewStore(testDB)

	user1 := createRandomUser(t)
	orderType1 := createRandomOrderType(t)
	discount1 := 1.0

	//run n concurrent order
	n := 5

	errs := make(chan error)
	results := make(chan OrderTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.OrderTx(context.Background(), OrderTxParams{
				UserName:    user1.Username,
				OrderTypeID: orderType1.ID,
				Discount:    discount1,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check order
		order := result.Order
		require.NotEmpty(t, order)
		require.Equal(t, order.UserName, user1.Username)
		require.Equal(t, order.OrderTypeID, orderType1.ID)
		require.Equal(t, order.Discount, discount1)
		require.NotZero(t, order.ID)
		require.NotZero(t, order.CreatedAt)

		_, err = store.GetOrder(context.Background(), order.ID)
		require.NoError(t, err)

		// check user
		user := result.User
		require.NotEmpty(t, user)
		require.Equal(t, user.Username, user1.Username)
		require.Equal(t, user.Email, user1.Email)
		require.Equal(t, user.HashedPassword, user1.HashedPassword)
	}
}
