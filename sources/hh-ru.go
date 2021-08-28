package sources

import (
	"encoding/json"
	"github.com/semyon-dev/gpn-tc-backend/model"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseHH(text string) (companies []model.HHItem, err error) {
	companies = []model.HHItem{}
	url := "https://api.hh.ru/employers"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return companies, err
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
		return companies, err
	}
	var hhReply model.HHReply
	err = json.Unmarshal(body, &hhReply)
	if err != nil {
		log.Println(err)
		return companies, err
	}
	return hhReply.Items, err
}
