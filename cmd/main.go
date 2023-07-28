package main

import (
	"encoding/json"
	"os"

	"github.com/gocolly/colly"
)

type Product struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageUrl string `json:"imageUrl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.jumia.com.ng"),
	)

	var products []Product

	c.OnHTML("a.core", func(h *colly.HTMLElement) {
		product := Product{
			Name: h.ChildText("div.name"),
			Price: h.ChildText("div.prc"),
			ImageUrl: h.ChildAttr("img", "data-src"),
		}
		products = append(products, product)
	})
	
	c.Visit("https://www.jumia.com.ng")
	


	content, err := json.Marshal(products)
	if err != nil {
		panic(err.Error())
	}

	os.WriteFile("content.json", content, 0644)

}
