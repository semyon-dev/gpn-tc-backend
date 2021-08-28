package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UtilityModel struct {
	Id                     primitive.ObjectID `bson:"_id" json:"id"`
	RegistrationNumber     string             `json:"registration number" bson:"registration number"`
	RegistrationDate       string             `json:"registration date" bson:"registration date"`
	ApplicationNumber      string             `json:"application number" bson:"application number"`
	ApplicationDate        string             `json:"application date" bson:"application date"`
	Authors                string             `json:"authors" bson:"authors"`
	AuthorsInLatin         string             `json:"authors in latin" bson:"authors in latin"`
	PatentHolders          string             `json:"patent holders" bson:"patent holders"`
	PatentHoldersInLatin   string             `json:"patent holders in latin" bson:"patent holders in latin"`
	UtilityModelName       string             `json:"utility model name" bson:"utility model name"`
	PatentStartingDate     string             `json:"patent starting date" bson:"patent starting date"`
	PatentGrantPublishDate string             `json:"patent grant publish date" bson:"patent grant publish date"`
	Actual                 string             `json:"actual" bson:"actual"`
	PublicationURL         string             `json:"publication URL" bson:"publication URL"`
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

type HHCompany struct {
	Name               string `json:"name"`
	Type               string `json:"type"`
	Id                 string `json:"id"`
	SiteUrl            string `json:"site_url"`
	Description        string `json:"description"`
	BrandedDescription string `json:"-"`
	VacanciesUrl       string `json:"vacancies_url"`
	OpenVacancies      int    `json:"open_vacancies"`
	Trusted            bool   `json:"trusted"`
	AlternateUrl       string `json:"alternate_url"`
	InsiderInterviews  []struct {
		Url   string `json:"url"`
		Id    string `json:"id"`
		Title string `json:"title"`
	} `json:"insider_interviews"`
	LogoUrls struct {
		Field1   string `json:"90"`
		Field2   string `json:"240"`
		Original string `json:"original"`
	} `json:"logo_urls"`
	Area struct {
		Url  string `json:"url"`
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"area"`
	Relations  []interface{} `json:"relations"`
	Industries []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"industries"`
}

type Suppliers struct {
	Companies []struct {
		AllSees string `json:"allSees"`
		City    string `json:"city"`
		ID      string `json:"id"`
		Name    string `json:"name"`
		Type    string `json:"type"`
	} `json:"companies"`
}

type RBK struct {
	Companies []struct {
		Category string `json:"category"`
		Link     string `json:"link"`
		Name     string `json:"name"`
		Text     string `json:"text"`
	} `json:"companies"`
}
