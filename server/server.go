package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server struct to store the gin Engine
type Server struct {
	Router *gin.Engine
}

// Run the server on dedicated port
func (server *Server) Run(port string) {
	gin.SetMode(gin.ReleaseMode)
	server.Router = gin.New()

	server.initRoutes()

	// Load HTML and Static files
	server.Router.LoadHTMLGlob("client/*.html")
	server.Router.Static("/css", "client/css")
	server.Router.Static("/img", "client/img")
	server.Router.Static("/js", "client/js")

	fmt.Printf(`SAC KGEC website live on port :%s!`, port)
	fmt.Println()
	server.Router.Run(":" + port)
}
