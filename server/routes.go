package server

var blogs = GetBlogs()

func (server *Server) initRoutes() {
	server.Router.GET("/", RenderHome)
	server.Router.GET("/faculty", RenderFaculty)
	for i := 0; i < len(blogs.Blogs); i++ {
		endpoint := "/" + blogs.Blogs[i].URL
		server.Router.GET(endpoint, RenderBlog(blogs.Blogs[i].Title, blogs.Blogs[i].Content))
	}
}
