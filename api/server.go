package api

import (
	"github.com/gin-gonic/gin"
	db "simplebank/db/sqlc"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

