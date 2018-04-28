package documentdb

import "fmt"
import "strings"
import "time"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "bytes"

// Client represents a client for interacting with the Comsmos DB REST API
type Client struct {
	*http.Client
	key string
	url string
}

// NewClient creates a new client using the provided master key and endpoint
func NewClient(key, url string) *Client {
	return &Client{
		Client: &http.Client{},
		key:    key,
		url:    url,
	}
}

// GetDatabase gets the database with the given id
func (client *Client) GetDatabase(id string) (*Database, error) {
	url := fmt.Sprintf("%v/dbs/%v", client.url, id)
	date := currentRFC1123FormattedDate(time.Now().UTC())

	provider := &MasterKeyTokenProvider{
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

	if res.StatusCode == 404 {
		return nil, nil
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unsupported response status code: %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	db := &Database{}

	err = json.Unmarshal(b, db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// ListDatabases gets all the databases for the Cosmos DB account
func (client *Client) ListDatabases() ([]*Database, error) {
	url := fmt.Sprintf("%v/dbs", client.url)
	date := currentRFC1123FormattedDate(time.Now().UTC())

	provider := &MasterKeyTokenProvider{
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
		return []*Database{}, err
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return []*Database{}, err
	}

	req.Header.Set("Authorization", signature)
	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2015-08-06")

	res, err := client.Do(req)

	if err != nil {
		return []*Database{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unsupported response status code: %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return []*Database{}, err
	}

	temp := &struct {
		Databases []*Database
	}{}

	err = json.Unmarshal(b, &temp)

	if err != nil {
		return []*Database{}, err
	}

	return temp.Databases, nil
}

// DeleteDatabase deletes the database with the given id
func (client *Client) DeleteDatabase(id string) error {
	url := fmt.Sprintf("%v/dbs/%v", client.url, id)
	date := currentRFC1123FormattedDate(time.Now().UTC())

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

	defer res.Body.Close()

	if res.StatusCode == 404 {
		return fmt.Errorf("Unable to delete database %v, database does not exist", id)
	}

	if res.StatusCode != 204 {
		return fmt.Errorf("Unsupported response status code: %v", res.StatusCode)
	}

	return nil
}

// CreateDatabase creates a database with the given id
func (client *Client) CreateDatabase(id string) (*Database, error) {
	url := fmt.Sprintf("%v/dbs", client.url)
	date := currentRFC1123FormattedDate(time.Now().UTC())

	provider := &MasterKeyTokenProvider{
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
		return nil, err
	}

	body := &struct {
		ID string `json:"id"`
	}{
		ID: id,
	}

	b, err := json.Marshal(&body)

	br := bytes.NewReader(b)

	req, err := http.NewRequest("POST", url, br)

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

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("Unable to delete database %v, database does not exist", id)
	}

	if res.StatusCode != 201 {
		return nil, fmt.Errorf("Unsupported response status code: %v", res.StatusCode)
	}

	b, err = ioutil.ReadAll(res.Body)

	db := &Database{}
	err = json.Unmarshal(b, db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func currentRFC1123FormattedDate(t time.Time) string {
	return strings.Replace(t.Format(time.RFC1123), "UTC", "GMT", 1)
}
