package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendBadRequest(c *gin.Context, err error) error {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return  err
}

func ArrayContains(array []string, target string) bool {
	for _, x := range array {
		if x == target {
			return true
		}
	}
	return false
}



