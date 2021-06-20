package server

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// RenderHome ...
func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title":  "Students' Automobile Club KGEC",
		"isHome": true,
	})
}

// RenderBlog ...
func RenderBlog(title, content string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "blog.html", gin.H{
			"title":   title,
			"content": template.HTML(content),
		})
	}
}

// RenderFaculty ...
func RenderFaculty(c *gin.Context) {
	c.HTML(200, "faculty.html", gin.H{
		"title": "Faculty",
	})
}

//RenderAboutUs
func RenderAboutUs(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"title":  "About Us",
		"isHome": false,
	})
}
