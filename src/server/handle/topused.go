package handle

import (
	"github.com/gin-gonic/gin"
	"log"
	"most-used-word/src/model"
	"most-used-word/src/services"
	"net/http"
	"strconv"
)

// handler for request get top used word from text
func HandleTopUsed(c *gin.Context) {

	var bodyContent *model.ContentReceive

	// get struct from body
	err := c.BindJSON(&bodyContent)
	if err != nil {
		log.Printf("%s", err.Error())
		responseError(c, err, http.StatusBadRequest)
		return
	}

	// convert string to int
	topNumber, err := strconv.Atoi(bodyContent.TopNumber)
	if err != nil {
		log.Printf("%s", err.Error())
		responseError(c, err, http.StatusInternalServerError)
		return
	}

	// get top used words from text and return
	topUse := services.TopUsedWord(bodyContent.Content, topNumber)
	result := model.TopUsedReturn{
		Title: "Return content",
		Result: topUse,
	}
	c.JSON(http.StatusOK, result)
}

// handle response error
func responseError(c *gin.Context, er error, statusCode int) {
	response := &model.ErrorResponse{
		Message:    er.Error(),
		StatusCode: statusCode,
	}
	c.JSON(statusCode, response)
	c.Abort()
}
