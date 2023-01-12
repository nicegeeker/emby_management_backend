package db

import (
	"context"
	"testing"

	"github.com/nicegeeker/emby_management_backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomOrderType(t *testing.T) OrderType {
	arg := CreateOrderTypeParams{
		Days:  util.RandomInt(0, 365),
		Price: util.RandomInt(1, 180),
	}

	orderType, err := testQueries.CreateOrderType(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, orderType)

	require.Equal(t, arg.Days, orderType.Days)
	require.Equal(t, arg.Price, orderType.Price)

	require.NotZero(t, orderType.ID)

	return orderType
}

func TestCreateOrderType(t *testing.T) {
	createRandomOrderType(t)
}
