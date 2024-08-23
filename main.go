package main

import (
	"fmt"
	"github.com/meilisearch/meilisearch-go"
)

func main() {
	//

	client := meilisearch.New("http://localhost:7700", meilisearch.WithAPIKey("password"))

	_, err := client.Index("products").UpdateFilterableAttributes(&[]string{
		"is_active",
		"selling_status",
	})
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%+v\n", resp1)

	resp, err := client.Index("products").Search("xin chao", &meilisearch.SearchRequest{
		Filter: "is_active = 1 AND selling_status = AVAILABLE",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp)
}
