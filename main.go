package main

import "log"
import "os"
import "fmt"

func main() {
	client := NewDocumentDBClient(os.Getenv("CosmosKey"), os.Getenv("CosmosUrl"))

	_, err := client.GetDatabase("scott-cdb-db")

	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.CreateDatabase("scott-cdb-db")

	if err != nil {
		log.Fatalln(err)
	}


	err = client.DeleteDatabase("scott-cdb-db")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("DONE")
}
