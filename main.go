package main

import (
	"fmt"
	"github.com/meilisearch/meilisearch-go"
)

func main() {
	//

	client := meilisearch.New("http://localhost:7700", meilisearch.WithAPIKey("password"))

	resp, err := client.Index("products").Search("code", &meilisearch.SearchRequest{
		Filter: "rating.users >= 90",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
