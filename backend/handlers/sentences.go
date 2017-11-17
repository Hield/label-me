package handlers

import (
	"log"
	"strconv"

	"github.com/Hield/label-me/backend/models"
	"github.com/gin-gonic/gin"
)

// AddSentence is the endpoint to add sentence
func AddSentence(c *gin.Context) {
	var err error
	var request struct {
		Content string `json:"content"`
		Label   string `json:"label"`
	}
	err = c.BindJSON(&request)
	if err != nil {
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
