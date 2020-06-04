package server

import (
	"server/middleware"

	"github.com/gin-gonic/gin"
)

// Server used to manage web
type Server struct {
	router   *gin.Engine
	listenIP string
}

// NewServer returns a server instance
func NewServer(listenIP string) *Server {
	server := &Server{
		router:   gin.New(),
		listenIP: listenIP,
	}

	server.router.Use(gin.Logger())
	server.router.Use(gin.Recovery())

	// Cors
	server.router.Use(middleware.Cors())
	return server
}

// Start used to start a http service
func (server *Server) Start() {
	server.router.Run(server.listenIP)
}

func (server *Server) Router() *gin.Engine {
	return server.router
}

// RegisterRouter used to add router to server
func (server *Server) RegisterRouter(method, url string, handler ...gin.HandlerFunc) {
	switch method {
	case "GET":
		server.router.GET(url, handler...)
	case "POST":
		server.router.POST(url, handler...)
	case "DELETE":
		server.router.DELETE(url, handler...)
	}
}
