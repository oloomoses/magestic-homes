package main

import (
	"fmt"
	"log"

	"githum.com/oloomoses/magestic-homes/internal/fetcher"
	"githum.com/oloomoses/magestic-homes/internal/parser"
)

func main() {
	url := "https://mansfield.craigslist.org/search/apa#search=2~gallery~0"
	fmt.Println("Fetcing data")
	htmlData, _ := fetcher.Fetch(url)

	// fmt.Println(htmlData)

	fmt.Println("Parsing html data")
	houseData, err := parser.ExtractItem(htmlData)

	if err != nil {
		log.Fatal("Parse Error", err)
	}

	fmt.Println(houseData)

}
