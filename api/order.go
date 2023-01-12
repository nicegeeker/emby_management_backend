package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/nicegeeker/emby_management_backend/db/sqlc"
	"github.com/nicegeeker/emby_management_backend/token"
)

type orderRequest struct {
	UserName    string  `json:"userName" binding:"required, alphanum"`
	OrderTypeID int64   `json:"orderTypeID" binding:"required"`
	Discount    float64 `json:"discount"`
}

type orderResponse struct {
	Order db.Order     `json:"order"`
	User  userResponse `json:"user"`
}

func (server *Server) order(ctx *gin.Context) {
	var req orderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.OrderTxParams{
		UserName:    authPayload.Username,
		OrderTypeID: req.OrderTypeID,
		Discount:    req.Discount,
	}

	result, err := server.store.OrderTx(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := orderResponse{
		Order: result.Order,
		User:  newUserResponse(result.User),
	}

	ctx.JSON(http.StatusOK, rsp)
}
