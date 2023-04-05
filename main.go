package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"github.com/gocolly/colly"
)

type Pokemons struct{
	url, image, name, price string
}
func main(){

	pokemon := Pokemons{}
	var pokemons []Pokemons
	c:=colly.NewCollector()
	c.OnHTML("li.product",func(h *colly.HTMLElement) {
		pokemon.url = h.ChildAttr("a","href")
		pokemon.image = h.ChildAttr("img","src")
		pokemon.name = h.ChildText("h2")
		pokemon.price = h.ChildText(".price")
		pokemons=append(pokemons, pokemon)
		fmt.Println(pokemons)
	})
	c.Visit("https://scrapeme.live/shop/")
	fmt.Println(c)
	
	file, err := os.Create("pokemons.csv")
	if err != nil{
		log.Fatalln("Failed to create ouput CSV file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	writer.Write(headers)
	for _, pokemon := range pokemons {
		record := []string{
			pokemon.url,
			pokemon.image,
			pokemon.name,
			pokemon.price,
		}
		writer.Write(record)
	}
	defer writer.Flush()
}
