package handlers

import (
	"errors"
	"log"
	"strconv"

	"github.com/Hield/label-me/backend/models"
	"github.com/gin-gonic/gin"
)

// AddSentence is the endpoint for POST request /api/sentences
func AddSentence(c *gin.Context) {
	var err error
	var request struct {
		Content string `json:"content"`
		Label   string `json:"label"`
	}
	err = c.BindJSON(&request)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	if request.Content == "" {
		c.AbortWithError(500, errors.New("Empty content"))
		return
	}
	label, err := strconv.Atoi(request.Label)
	if err != nil || label > 2 || label < -2 {
		err = models.AddSentence(request.Content, false, 0)
	} else {
		err = models.AddSentence(request.Content, true, label)
	}
	if err != nil {
		log.Printf("Error: %s\n", err)
		c.AbortWithError(500, err)
		return
	}
}

// GetSentences is the endpoint for GET request /api/sentences
func GetSentences(c *gin.Context) {
	var err error
	nString := c.Query("n")
	n, err := strconv.Atoi(nString)
	if err != nil {
		c.AbortWithError(500, errors.New("Invalid number of sentences"))
		return
	}
	sentences, err := models.GetRandomSentences(n)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, sentences)
}

// GetMostPositiveSentences handle GET request to /api/sentences/positive
func GetMostPositiveSentences(c *gin.Context) {
	sentences, err := models.GetMostPositiveSentences()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, sentences)
}
