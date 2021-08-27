package sources

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type HHReply struct {
	PerPage int    `json:"per_page"`
	Page    int    `json:"page"`
	Pages   int    `json:"pages"`
	Found   int    `json:"found"`
	Items   []Item `json:"items"`
}

type Item struct {
	Id            string      `json:"id"`
	Name          string      `json:"name"`
	Url           string      `json:"url"`
	AlternateUrl  string      `json:"alternate_url"`
	VacanciesUrl  string      `json:"vacancies_url"`
	OpenVacancies int         `json:"open_vacancies"`
	LogoUrls      interface{} `json:"logo_urls"`
}

func ParseHH(text string) (companies []Item, err error) {
	url := "https://api.hh.ru/employers"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("User-Agent", "gpn-tc-backend/1.0 (semennovikov1@yandex.ru)")
	req.Header.Set("HH-User-Agent", "gpn-tc-backend/1.0 (semennovikov1@yandex.ru)")
	q := req.URL.Query()
	q.Add("text", text)
	req.URL.RawQuery = q.Encode()
	// Отправляем запрос
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var hhReply HHReply
	err = json.Unmarshal(body, &hhReply)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return hhReply.Items, err
}
