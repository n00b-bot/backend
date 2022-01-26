package api

import (
	db "backend/db/sqlc"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64   `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64   `json:"to_account_id" binding:"required,min=1"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
	Currency      string  `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	if !server.validAccount(ctx, req.FromAccountID, req.Currency) {
		return
	}
	if !server.validAccount(ctx, req.ToAccountID, req.Currency) {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}
	transfer, err := server.store.TransferTx(ctx, db.TransferTxParams(arg))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transfer)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return false
	}
	if account.Currency != currency {
		ctx.JSON(http.StatusBadRequest, errResponse(fmt.Errorf("not match currency")))
		return false
	}
	return true
}
