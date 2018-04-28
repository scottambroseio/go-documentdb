package main

import "fmt"
import "strings"
import "time"
import "net/http"
import "io/ioutil"
import "encoding/json"

type documentDBClient struct {
	*http.Client
	key string
	url string
}

func NewDocumentDBClient(key, url string) *documentDBClient {
	return &documentDBClient{
		Client: &http.Client{},
		key:    key,
		url:    url,
	}
}

func currentRFC1123FormattedDate(t time.Time) string {
	return strings.Replace(t.Format(time.RFC1123), "UTC", "GMT", 1)

}

func (client *documentDBClient) DeleteDatabase(id string) error {
	url := fmt.Sprintf("%v/dbs/%v", client.url, id)
	date := currentRFC1123FormattedDate(time.Now())

	provider := &MasterKeyTokenProvider{
		Verb:         "DELETE",
		ResourceType: "dbs",
		ResourceLink: fmt.Sprintf("dbs/%v", id),
		Date:         date,
		Key:          client.key,
		KeyType:      "master",
		TokenVersion: "1.0",
	}

	signature, err := provider.GenerateToken()

	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", signature)
	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2015-08-06")

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	res.Body.Close()
	return nil
}

func (client *documentDBClient) CreateDatabase(id string) (string, error) {
	url := fmt.Sprintf("%v/dbs", client.url)
	date := currentRFC1123FormattedDate(time.Now())

	provider := MasterKeyTokenProvider{
		Verb:         "POST",
		ResourceType: "dbs",
		ResourceLink: "",
		Date:         date,
		Key:          client.key,
		KeyType:      "master",
		TokenVersion: "1.0",
	}

	signature, err := provider.GenerateToken()

	if err != nil {
		return "", err
	}

	body := strings.NewReader(`{ "id": "` + id + `" }`)
	req, err := http.NewRequest("POST", url, body)

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", signature)
	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2015-08-06")

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (client *documentDBClient) GetDatabase(id string) (*Database, error) {
	url := fmt.Sprintf("%v/dbs/%v", client.url, id)
	date := currentRFC1123FormattedDate(time.Now())

	provider := MasterKeyTokenProvider{
		Verb:         "GET",
		ResourceType: "dbs",
		ResourceLink: fmt.Sprintf("dbs/%v", id),
		Date:         date,
		Key:          client.key,
		KeyType:      "master",
		TokenVersion: "1.0",
	}

	signature, err := provider.GenerateToken()

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", signature)
	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2015-08-06")

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	resp := &Database{}

	err = json.Unmarshal(bytes, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (client *documentDBClient) ListDatabases() (string, error) {
	url := fmt.Sprintf("%v/dbs", client.url)
	date := currentRFC1123FormattedDate(time.Now())

	provider := MasterKeyTokenProvider{
		Verb:         "GET",
		ResourceType: "dbs",
		ResourceLink: "",
		Date:         date,
		Key:          client.key,
		KeyType:      "master",
		TokenVersion: "1.0",
	}

	signature, err := provider.GenerateToken()

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", signature)
	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2015-08-06")

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
