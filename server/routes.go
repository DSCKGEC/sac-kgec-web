package server

func (server *Server) initRoutes() {
	server.Router.GET("/", RenderHome)
}
