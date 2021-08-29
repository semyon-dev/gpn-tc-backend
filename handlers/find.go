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
		Type string `json:"type"`
	}{}
	err := c.ShouldBindJSON(&jsonInput)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad input json"})
		return
	}

	switch jsonInput.Type {
	case "block":
		req := strings.Split(jsonInput.Text, " ")
		for _, v := range req {
			if v == "OR" {

			} else if v == "AND" {

			} else if len(v) == 2 && strings.Contains(v, "W") {

			} else if len(v) == 2 && strings.Contains(v, "D") {

			} else if strings.Contains(v, "+") {

			}
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	case "all":
		err, res := FindByNameAll(jsonInput.Text)
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"hhRu":       res[0].([]model.HHCompany),
			"rospatent":  res[1].([]model.UtilityModel),
			"habrCareer": res[2].(model.HabrCareer).Companies,
			"suppliers":  res[3].(model.Suppliers).Companies,
			"RBC":        res[4].(model.RBK).Companies,
		})
	case "":
		err, res := FindByNameAll(jsonInput.Text)
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"hhRu":       res[0].([]model.HHCompany),
			"rospatent":  res[1].([]model.UtilityModel),
			"habrCareer": res[2].(model.HabrCareer).Companies,
			"suppliers":  res[3].(model.Suppliers).Companies,
			"RBC":        res[4].(model.RBK).Companies,
		})
	case "company":
	case "rospatent":
		res := FindRosPatent(jsonInput.Text)
		c.JSON(http.StatusOK, gin.H{
			"rospatent": res,
			"message":   "ok",
		})
	case "okved":
		res := db.FindOkved(jsonInput.Text)
		if len(res) == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "no result",
			})
			return
		}
		comps, err := sources.ParseOkved(strings.ReplaceAll(res[0].Link, "/contragent/by-okved/", ""))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"okved":   comps.Companies,
			"message": "ok",
		})
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid type",
		})
	}
}

func FindRosPatent(text string) []model.UtilityModel {
	itemsRospatent := db.FindInUtilityModel(text)
	for i, v := range itemsRospatent {
		if len(v.PatentHolders) == 0 {
			itemsRospatent[i].Bench = Bench(v.UtilityModelName, text, 50)
			continue
		}
		itemsRospatent[i].Bench = Bench(v.PatentHolders, text, 100)
	}
	sort.Slice(itemsRospatent, func(i, j int) bool {
		return itemsRospatent[i].Bench > itemsRospatent[j].Bench
	})
	return itemsRospatent
}

func FindByNameAll(text string) (err error, res []interface{}) {
	var wg sync.WaitGroup
	wg.Add(5)
	var itemsHH []model.HHCompany
	var itemsRospatent []model.UtilityModel
	var parseHabr model.HabrCareer
	var parseSuppliers model.Suppliers
	var parseRBK model.RBK
	go func() {
		itemsHH, err = sources.ParseHH(text)
		for i, v := range itemsHH {
			if len(v.Description) == 0 {
				itemsHH[i].Bench = Bench(v.Name, text, 50)
				continue
			}
			itemsHH[i].Bench = Bench(v.Description, text, 100)
		}
		sort.Slice(itemsHH, func(i, j int) bool {
			return itemsHH[i].Bench > itemsHH[j].Bench
		})
		wg.Done()
	}()
	go func() {
		itemsRospatent = FindRosPatent(text)
		wg.Done()
	}()
	go func() {
		parseHabr, err = sources.ParseHabr(text)
		for i, v := range parseHabr.Companies {
			if len(v.Description) == 0 {
				parseHabr.Companies[i].Bench = Bench(v.Name, text, 50)
				continue
			}
			parseHabr.Companies[i].Bench = Bench(strings.Join(v.Description, ""), text, 100)
		}
		sort.Slice(parseHabr.Companies, func(i, j int) bool {
			return parseHabr.Companies[i].Bench > parseHabr.Companies[j].Bench
		})
		wg.Done()
	}()
	go func() {
		parseSuppliers, err = sources.ParseSuppliers(text)
		for i, v := range parseSuppliers.Companies {
			if len(v.Type) == 0 {
				parseSuppliers.Companies[i].Bench = Bench(v.Name, text, 50)
				continue
			}
			parseSuppliers.Companies[i].Bench = Bench(v.Type, text, 100)
		}
		sort.Slice(parseSuppliers.Companies, func(i, j int) bool {
			return parseSuppliers.Companies[i].Bench > parseSuppliers.Companies[j].Bench
		})
		wg.Done()
	}()
	go func() {
		parseRBK, err = sources.ParseRBK(text)
		for i, v := range parseRBK.Companies {
			if len(v.Text) == 0 {
				itemsRospatent[i].Bench = Bench(v.Name, text, 50)
				continue
			}
			parseRBK.Companies[i].Bench = Bench(v.Text, text, 100)
		}
		sort.Slice(parseRBK.Companies, func(i, j int) bool {
			return parseRBK.Companies[i].Bench > parseRBK.Companies[j].Bench
		})
		wg.Done()
	}()
	wg.Wait()
	res = append(res, itemsHH, itemsRospatent, parseHabr, parseSuppliers, parseRBK)
	return err, res
}

func Bench(text1, text2 string, k float64) int {
	if text1 == text2 {
		return 99
	}
	bench := util.DistanceForStrings([]rune(strings.ToLower(text1)), []rune(strings.ToLower(text2)))
	return int(math.Round(float64(bench) / float64(len(text1)) * k))
}
