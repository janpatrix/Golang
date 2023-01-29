package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Title string
	Body  string
}

var posts []Post

func main() {
	router := gin.Default()

	// initialize the posts slice
	posts = []Post{
		{Title: "First post", Body: "This is the first post on the blog."},
		{Title: "Second post", Body: "This is the second post on the blog."},
	}

	// handle requests to the root URL
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"posts": posts,
		})
	})

	// handle requests to the "post" URL
	router.GET("/post/:id", func(c *gin.Context) {
		id := c.Param("id")
		post := posts[id]
		c.HTML(http.StatusOK, "post.html", gin.H{
			"post": post,
		})
	})

	// start the server
	router.Run(":8080")
}
