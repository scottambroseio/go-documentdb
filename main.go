package main

import "log"
import "fmt"
import "strings"
import "time"
import "net/http"
import "io/ioutil"
import "os"

func DeleteDatabase(id string) string {
	key := os.Getenv("CosmosKey")
	url := os.Getenv("CosmosUrl") + fmt.Sprintf("/dbs/%v", id)

	date := strings.Replace(time.Now().Format(time.RFC1123), "UTC", "GMT", 1)

	provider := MasterKeyTokenProvider{
		Verb:         "DELETE",
		ResourceType: "dbs",
		ResourceLink: fmt.Sprintf("dbs/%v", id),
		Date:         date,
		Key:          key,
		KeyType:      "master",
		TokenVersion: "1.0",
	}

	signature, err := provider.GenerateToken()

	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", signature)
	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2015-08-06")

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(bytes)
}

func CreateDatabase(id string) string {
	key := os.Getenv("CosmosKey")
	url := os.Getenv("CosmosUrl") + "/dbs"

	date := strings.Replace(time.Now().Format(time.RFC1123), "UTC", "GMT", 1)

	provider := MasterKeyTokenProvider{
		Verb:         "POST",
		ResourceType: "dbs",
		ResourceLink: "",
		Date:         date,
		Key:          key,
		KeyType:      "master",
		TokenVersion: "1.0",
	}

	signature, err := provider.GenerateToken()

	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	body := strings.NewReader(`{ "id": "` + id + `" }`)
	req, err := http.NewRequest("POST", url, body)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", signature)
	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2015-08-06")

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(bytes)
}

func GetDatabases() string {
	key := os.Getenv("CosmosKey")
	url := os.Getenv("CosmosUrl") + "/dbs"
	date := strings.Replace(time.Now().Format(time.RFC1123), "UTC", "GMT", 1)

	provider := MasterKeyTokenProvider{
		Verb:         "GET",
		ResourceType: "dbs",
		ResourceLink: "",
		Date:         date,
		Key:          key,
		KeyType:      "master",
		TokenVersion: "1.0",
	}

	signature, err := provider.GenerateToken()

	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", signature)
	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2015-08-06")

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(bytes)
}

func main() {
}
