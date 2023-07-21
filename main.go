package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"web-scrapper/models"

	"github.com/gocolly/colly"
)

func main() {
	allQuotes := make([]models.Quotes, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	collector.OnHTML(".quote", func(element *colly.HTMLElement) {
		quote := element.ChildText(".text")
		author := element.ChildText(".author")

		Q := models.Quotes{
			Author: &author,
			Quote:  &quote,
		}
		allQuotes = append(allQuotes, Q)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting", request.URL.String())
	})

	collector.Visit("http://quotes.toscrape.com")

	writeJSON(allQuotes)
}

func writeJSON(data []models.Quotes) {
	file, err := os.Create("Quotes.json")
	if err != nil {
		log.Println("unable to create json file")
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(data); err != nil {
		log.Println("unable to write to json file")
		return
	}
}
