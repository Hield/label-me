package handlers

import (
	"errors"
	"strconv"

	"github.com/Hield/label-me/backend/models"
	"github.com/gin-gonic/gin"
)

// AddAnswer handle POST request to /api/answers
func AddAnswer(c *gin.Context) {
	var err error
	var request struct {
		SentenceID string `json:"sentenceID"`
		Label      string `json:"label"`
		Feedback   string `json:"feedback"`
		IsSkipped  bool   `json:"isSkipped"`
	}
	err = c.BindJSON(&request)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	sentenceID, err := strconv.Atoi(request.SentenceID)
	if err != nil {
		c.AbortWithError(500, errors.New("Invalid sentence id"))
		return
	}

	var label int
	if !request.IsSkipped {
		label, err = strconv.Atoi(request.Label)
		if err != nil || label > 2 || label < -2 {
			c.AbortWithError(500, errors.New("Invalid label"))
			return
		}
	}
	err = models.AddAnswer(sentenceID, label, request.Feedback, request.IsSkipped)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
}
