package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/semyon-dev/gpn-tc-backend/sources"
	"log"
	"net/http"
)

func Find(c *gin.Context) {
	jsonInput := struct {
		Name string `json:"name"`
	}{}
	err := c.ShouldBindJSON(&jsonInput)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad req"})
		return
	}
	items, err := sources.ParseHH(jsonInput.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"items":   items,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"items":   items,
	})
}
