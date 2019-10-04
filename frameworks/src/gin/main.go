package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

//Note struct for notes
type Note struct {
	Creator string `json:"creator"`
	Text    string `json:"text" binding:"required"`
}

var notes = []Note{Note{Creator: "Auto", Text: "Exampole"}, Note{Creator: "Auto", Text: "Example2"}}

func main() {
	fmt.Println("Starting")
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/notes", NotesHandler)
		api.POST("/notes", NotesHandler)
	}
	router.Run(":3000")
}

//NotesHandler will handle Note Requests
func NotesHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if c.Request.Method == http.MethodGet {
		c.JSON(http.StatusOK, notes)
	}
}
