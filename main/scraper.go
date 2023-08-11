package main

import (
	"fmt"
	"os"
	"strings"

	// import Colly
	"github.com/gocolly/colly"
)

type Dictionary map[string]interface{}

type RecipeSpecs struct {
	difficulty, prepTime, cookingTime, servingSize, priceTier string
}

type Recipe struct {
	url, name      string
	ingredients    []Dictionary
	specifications RecipeSpecs
}

func main() {
	args := os.Args
	url := args[1]
	c := colly.NewCollector()
	// initializing the slice of structs that will contain the scraped data
	var recipes []Recipe

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Blimey, an error occurred!:", e)
	})
	c.OnHTML("main", func(main *colly.HTMLElement) {
		recipe := Recipe{}
		ingredients_dictionary := Dictionary{}
		recipe.url = url
		// find the recipe title and assign it to the struct
		recipe.name = main.ChildText(".gz-title-recipe")
		println("Scraping recipe for:", recipe.name)

		// iterate over each instance of a recipe spec
		// and assign its value to the recipe spec struct and the recipe
		main.ForEach(".gz-name-featured-data", func(i int, specListElement *colly.HTMLElement) {
			if strings.Contains(specListElement.Text, "Difficoltà: ") {
				recipe.specifications.difficulty = specListElement.ChildText("strong")
			}
			if strings.Contains(specListElement.Text, "Preparazione: ") {
				recipe.specifications.prepTime = specListElement.ChildText("strong")
			}
			if strings.Contains(specListElement.Text, "Cottura: ") {
				recipe.specifications.cookingTime = specListElement.ChildText("strong")
			}
			if strings.Contains(specListElement.Text, "Dosi per: ") {
				recipe.specifications.servingSize = specListElement.ChildText("strong")
			}
			if strings.Contains(specListElement.Text, "Costo: ") {
				recipe.specifications.priceTier = specListElement.ChildText("strong")
			}
		})

		// find the recipe introduction and ingredients and assign to the struct
		main.ForEach(".gz-ingredient", func(i int, ingredient *colly.HTMLElement) {
			ingredients_dictionary[ingredient.ChildText("a")] = ingredient.ChildText("span")
		})
		recipe.ingredients = append(recipe.ingredients, ingredients_dictionary)
		recipes = append(recipes, recipe)
	})

	c.Visit(url)

	fmt.Println(recipes)

}
