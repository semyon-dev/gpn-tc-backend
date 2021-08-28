package sources

import (
	"encoding/json"
	"fmt"
	"github.com/semyon-dev/gpn-tc-backend/model"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func ParseHH(text string) (companies []model.HHCompany, err error) {
	companies = []model.HHCompany{}
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
	var companiesAdds []model.HHCompany
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i, v := range hhReply.Items {
		if i == 3 {
			break
		}
		wg.Add(1)
		go func(url string) {
			comp, _ := ParseHHAdditional(url)
			mutex.Lock()
			companiesAdds = append(companiesAdds, comp)
			mutex.Unlock()
			wg.Done()
		}(v.Url)
	}
	wg.Wait()
	return companiesAdds, err
}

func ParseHHAdditional(url string) (company model.HHCompany, err error) {
	fmt.Println(url)
	company = model.HHCompany{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return company, err
	}
	req.Header.Set("User-Agent", "gpn-tc-backend/1.0 (semennovikov1@yandex.ru)")
	req.Header.Set("HH-User-Agent", "gpn-tc-backend/1.0 (semennovikov1@yandex.ru)")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return company, err
	}
	err = json.Unmarshal(body, &company)
	if err != nil {
		log.Println(err)
		return company, err
	}
	return company, err
}
