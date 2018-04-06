package main

import "log"
import "fmt"
import "os"

func main() {
	client := NewDocumentDBClient(os.Getenv("CosmosKey"), os.Getenv("CosmosUrl"))

	resp, err := client.GetDatabases()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
	fmt.Println()

	resp, err = client.CreateDatabase("scott-cdb-db")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
	fmt.Println()

	resp, err = client.GetDatabases()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
	fmt.Println()

	resp, err = client.DeleteDatabase("scott-cdb-db")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
	fmt.Println()

	resp, err = client.GetDatabases()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
	fmt.Println()
}
