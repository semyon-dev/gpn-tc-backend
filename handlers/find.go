package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/semyon-dev/gpn-tc-backend/db"
	"github.com/semyon-dev/gpn-tc-backend/model"
	"github.com/semyon-dev/gpn-tc-backend/sources"
	"log"
	"net/http"
	"sync"
)

func Find(c *gin.Context) {
	jsonInput := struct {
		Text string `json:"text"`
	}{}
	err := c.ShouldBindJSON(&jsonInput)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad req"})
		return
	}
	var wg sync.WaitGroup
	wg.Add(5)
	var itemsHH []model.HHCompany
	var itemsRospatent []model.UtilityModel
	var parseHabr model.HabrCareer
	var parseSuppliers model.Suppliers
	var parseRBK model.RBK
	go func() {
		itemsHH, err = sources.ParseHH(jsonInput.Text)
		wg.Done()
	}()
	go func() {
		itemsRospatent = db.FindInUtilityModel(jsonInput.Text)
		wg.Done()
	}()
	go func() {
		parseHabr, err = sources.ParseHabr(jsonInput.Text)
		wg.Done()
	}()
	go func() {
		parseSuppliers, err = sources.ParseSuppliers(jsonInput.Text)
		wg.Done()
	}()
	go func() {
		parseRBK, err = sources.ParseRBK(jsonInput.Text)
		wg.Done()
	}()
	wg.Wait()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"hhRu":       itemsHH,
		"rospatent":  itemsRospatent,
		"habrCareer": parseHabr.Companies,
		"suppliers":  parseSuppliers.Companies,
		"RBC":        parseRBK.Companies,
	})
}
