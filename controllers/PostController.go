package controllers

import (
	"goCurd/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

// post index
func (p *PostController) Index(c *gin.Context) {
	random_str := helpers.Random("uuid", 8)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Posts 1",
		"content": random_str,
	})
}

// create post
func (p *PostController) Create(c *gin.Context) {
	rand_str := helpers.Crypt("abc")
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "post create and rand string :" + rand_str,
	})
}

// post show
func (p *PostController) Show(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "post show , id: " + id,
	})
}

// post edit
func (p *PostController) Edit(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "post edit id: " + id,
	})
}

// post store
func (p *PostController) Store(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "post store id: " + id,
	})
}

// post delete
func (p *PostController) Delete(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "post delete id: " + id,
	})
}
