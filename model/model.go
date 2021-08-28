package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	Name        string
	Description string
	Authors     []interface{}
	Patents     []interface{}
	Contacts    interface{}
}

type UtilityModel struct {
	Id                     primitive.ObjectID `bson:"_id" json:"id"`
	RegistrationNumber     string             `json:"registration number"`
	RegistrationDate       string             `json:"registration date"`
	ApplicationNumber      string             `json:"application number"`
	ApplicationDate        string             `json:"application date"`
	Authors                string             `json:"authors"`
	AuthorsInLatin         string             `json:"authors in latin"`
	PatentHolders          string             `json:"patent holders"`
	PatentHoldersInLatin   string             `json:"patent holders in latin"`
	UtilityModelName       string             `json:"utility model name" bson:"utility model name"`
	PatentStartingDate     string             `json:"patent starting date"`
	PatentGrantPublishDate string             `json:"patent grant publish date"`
	Actual                 string             `json:"actual"`
	PublicationURL         string             `json:"publication URL"`
}

type HabrCareer struct {
	Companies []struct {
		Addresses []string `json:"addresses"`
		Contacts  []struct {
			Link  string `json:"link"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"contacts"`
		Description []string `json:"description"`
		Employees   []struct {
			Avatar      string `json:"avatar"`
			EmployeeUrl string `json:"employee_url"`
			Position    string `json:"position"`
			Username    string `json:"username"`
		} `json:"employees"`
		Logo   string   `json:"logo"`
		Name   string   `json:"name"`
		Site   string   `json:"site"`
		Skills []string `json:"skills"`
	} `json:"companies"`
}

type HHReply struct {
	PerPage int      `json:"per_page"`
	Page    int      `json:"page"`
	Pages   int      `json:"pages"`
	Found   int      `json:"found"`
	Items   []HHItem `json:"items"`
}

type HHItem struct {
	Id            string      `json:"id"`
	Name          string      `json:"name"`
	Url           string      `json:"url"`
	AlternateUrl  string      `json:"alternate_url"`
	VacanciesUrl  string      `json:"vacancies_url"`
	OpenVacancies int         `json:"open_vacancies"`
	LogoUrls      interface{} `json:"logo_urls"`
}