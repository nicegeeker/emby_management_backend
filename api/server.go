package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/nicegeeker/emby_management_backend/db/sqlc"
	"github.com/nicegeeker/emby_management_backend/token"
	"github.com/nicegeeker/emby_management_backend/util"
)

// 用于给所有http请求提供服务
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmeticKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/api/login", server.loginUser)
	router.POST("/api/register", server.createUser)
	router.POST("/api/refresh", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/api/user", server.getUser)
	authRoutes.POST("/api/list_users", server.listUsers)

	authRoutes.POST("/api/order", server.order)

	server.router = router
}

//用于启动服务
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
