package sources

import (
	"encoding/json"
	"github.com/semyon-dev/gpn-tc-backend/config"
	"github.com/semyon-dev/gpn-tc-backend/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ParseSuppliers(text string) (companies model.Suppliers, err error) {
	companies = model.Suppliers{}
	url := config.ParseSuppliers
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return companies, err
	}
	req.Header.Set("User-Agent", "gpn-tc-backend/1.0")
	q := req.URL.Query()
	q.Add("q", text)
	req.URL.RawQuery = q.Encode()
	// Отправляем запрос
	client := &http.Client{}
	client.Timeout = 16 * time.Second
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return companies, err
	}
	err = json.Unmarshal(body, &companies)
	if err != nil {
		log.Println(err)
		return companies, err
	}
	return companies, err
}
