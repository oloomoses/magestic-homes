package main

import (
	"fmt"
	"log"

	"githum.com/oloomoses/magestic-homes/internal/fetcher"
	"githum.com/oloomoses/magestic-homes/internal/parser"
	"githum.com/oloomoses/magestic-homes/internal/storage"
)

func main() {
	url := "https://mansfield.craigslist.org/search/apa#search=2~gallery~0"
	htmlData, _ := fetcher.Fetch(url)

	fmt.Println(htmlData)

	fmt.Println("Parsing html data")
	title, err := parser.ExtactTitle(`"<div class="title">2 bedroom, BBQ/Picnic Area, In Ashland</div>"`)

	if err != nil {
		log.Fatal("Parse Error", err)
	}

	storage.Print("Title: " + title)

	err = storage.SaveToFile("output.txt", title)
	if err != nil {
		log.Fatal("Save error: ", err)
	}

}
