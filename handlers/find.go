package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/semyon-dev/gpn-tc-backend/db"
	"github.com/semyon-dev/gpn-tc-backend/model"
	"github.com/semyon-dev/gpn-tc-backend/sources"
	"github.com/semyon-dev/gpn-tc-backend/util"
	"log"
	"math"
	"net/http"
	"sort"
	"strings"
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
		for i, v := range itemsHH {
			if len(v.Description) == 0 {
				itemsHH[i].Bench = Bench(v.Name, jsonInput.Text, 50)
				continue
			}
			itemsHH[i].Bench = Bench(v.Description, jsonInput.Text, 100)
		}
		sort.Slice(itemsHH, func(i, j int) bool {
			return itemsHH[i].Bench > itemsHH[j].Bench
		})
		wg.Done()
	}()
	go func() {
		itemsRospatent = db.FindInUtilityModel(jsonInput.Text)
		for i, v := range itemsRospatent {
			if len(v.PatentHolders) == 0 {
				itemsRospatent[i].Bench = Bench(v.UtilityModelName, jsonInput.Text, 50)
				continue
			}
			itemsRospatent[i].Bench = Bench(v.PatentHolders, jsonInput.Text, 100)
		}
		sort.Slice(itemsRospatent, func(i, j int) bool {
			return itemsRospatent[i].Bench > itemsRospatent[j].Bench
		})
		wg.Done()
	}()
	go func() {
		parseHabr, err = sources.ParseHabr(jsonInput.Text)
		for i, v := range parseHabr.Companies {
			if len(v.Description) == 0 {
				parseHabr.Companies[i].Bench = Bench(v.Name, jsonInput.Text, 50)
				continue
			}
			parseHabr.Companies[i].Bench = Bench(strings.Join(v.Description, ""), jsonInput.Text, 100)
		}
		sort.Slice(parseHabr.Companies, func(i, j int) bool {
			return parseHabr.Companies[i].Bench > parseHabr.Companies[j].Bench
		})
		wg.Done()
	}()
	go func() {
		parseSuppliers, err = sources.ParseSuppliers(jsonInput.Text)
		wg.Done()
	}()
	go func() {
		parseRBK, err = sources.ParseRBK(jsonInput.Text)
		for i, v := range parseRBK.Companies {
			if len(v.Text) == 0 {
				itemsRospatent[i].Bench = Bench(v.Name, jsonInput.Text, 50)
				continue
			}
			parseRBK.Companies[i].Bench = Bench(v.Text, jsonInput.Text, 100)
		}
		sort.Slice(parseRBK.Companies, func(i, j int) bool {
			return parseRBK.Companies[i].Bench > parseRBK.Companies[j].Bench
		})
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

func Bench(text1, text2 string, k float64) int {
	if text1 == text2 {
		return 99
	}
	bench := util.DistanceForStrings([]rune(strings.ToLower(text1)), []rune(strings.ToLower(text2)))
	return int(math.Round(float64(bench) / float64(len(text1)) * k))
}
