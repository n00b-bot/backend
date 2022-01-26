package api

import (
	db "backend/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Fullname string `json:"full_name" binding:"required"`
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username:     req.Username,
		HashedPasswd: req.Password,
		FullName:     req.Fullname,
		Email:        req.Email,
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}
