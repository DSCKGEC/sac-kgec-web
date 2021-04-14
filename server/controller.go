package server

import (
	"github.com/gin-gonic/gin"
)

// RenderHome ...
func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title":  "Students' Automobile Club KGEC",
		"isHome": true,
	})
}
