package main

import "log"
import "fmt"
import "crypto/hmac"
import "crypto/sha256"
import "encoding/base64"
import "net/url"
import "strings"
import "time"
import "net/http"
import "io/ioutil"

func main() {
	verb := "GET"
	resourceType := "dbs"
	resourceLink := ""
	date := strings.Replace(time.Now().Format(time.RFC1123), "UTC", "GMT", 1)
	key := ""
	keyType := "master"
	tokenVersion := "1.0"

	signature := GenerateMasterKeyAuthorizationSignature(verb, resourceType, resourceLink, key, keyType, tokenVersion, date)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "", nil)

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

	fmt.Println(string(bytes))
}

func GenerateMasterKeyAuthorizationSignature(verb, resourceType, resourceId, key, keyType, tokenVersion, date string) string {
	decoded, err := base64.StdEncoding.DecodeString(key)

	if err != nil {
		log.Fatalln(err)
	}

	mac := hmac.New(sha256.New, decoded)
	payload := fmt.Sprintf("%v\n%v\n%v\n%v\n\n", strings.ToLower(verb), strings.ToLower(resourceType), resourceId, strings.ToLower(date))
	mac.Write([]byte(payload))
	hashPayload := mac.Sum(nil)

	sig := base64.StdEncoding.EncodeToString(hashPayload)
	e := fmt.Sprintf("type=%v&ver=%v&sig=%v", keyType, tokenVersion, sig)
	return url.QueryEscape(e)
}
