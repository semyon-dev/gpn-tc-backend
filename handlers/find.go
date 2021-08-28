package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/semyon-dev/gpn-tc-backend/db"
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
	itemsHH, err := sources.ParseHH(jsonInput.Name)
	itemsRospatent := db.FindInUtilityModel(jsonInput.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "ok",
		"hh":          itemsHH,
		"rospatent":   itemsRospatent,
		"habr-career": itemsRospatent,
	})
}
