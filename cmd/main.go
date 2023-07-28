package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Product struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageUrl string `json:"imageUrl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("jumia.com.ng"),
	)

	fmt.Println("about to")

	c.OnHTML("a.core", func(h *colly.HTMLElement) {
		fmt.Println("=======")
		fmt.Println(h.Text)
	})

	c.Visit("https://www.jumia.com.ng")
}
