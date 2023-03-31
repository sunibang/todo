package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer() (*Server, error) {
	server := &Server{}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/todos", server.getTodos)
	router.POST("/todos", server.addTodo)
	router.GET("/todos/:id", server.getTodo)
	router.PATCH("/todos/:id", server.toggleTodoStatus)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
