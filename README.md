# Webscraper
A "GialloZafferano" web scraper command line application built in Go using Colly.
The app is designed to scrape the recipe pages of the GialloZafferano website.
This is not affiliated with GialloZafferano, I just like italian food. 

## Introduction 👣

In this tutorial, we will walk through the process of building a web scraper using the Colly library in the Go programming language. This scraper will extract recipe data from a specific website. We'll cover the code step by step to understand each component.

### What is a web-scraper

A web-scraper is, as the name suggests, a tool you can use to “scrape” or extract data from websites.  
For this project, I decided to scrape the Italian recipe website [Giallo Zafferano](https://www.giallozafferano.it/).

### Getting started

To get started

* 💻 your IDE of choice - I’m using VSCode.
    
* 🍪 some snacks because all this recipe scraping will make you hungry.
    

## Let's get coding 🚀

### Install Go

Firstly, let's set up our Go project. As this is the first article of the Go Learning Series, I'll spend a few seconds talking you through installing Go.  
If you have installed it already, then scroll through ⏭️

Depending on your operating system, you can find your installation guide on [Go documentation page](https://go.dev/doc/install). If you're a macOS user and you use `brew`, you can run it in your terminal:

```bash
brew install go
```

### Set up the project

Create a new directory for your project, move to the directory then run the following command, where you can replace the word `webscraper` with anything you want your module to be named.

```bash
go mod init webscraper
```

<div data-node-type="callout">
<div data-node-type="callout-emoji">💡</div>
<div data-node-type="callout-text">The <code>go mod init</code> command initializes a new Go module in the directory where it is executed. It creates a new <code>go.mod</code> file, which is used to define dependencies and manage versions of third-party packages used in the project (👀 kinda like a package.json if you're using node).</div>
</div>

Now let's install `colly` and its dependencies

```bash
go get github.com/gocolly/colly
```

This command will also update the `go.mod` file with all the required dependencies as well as create a `go.sum` file.

<div data-node-type="callout">
<div data-node-type="callout-emoji">💡</div>
<div data-node-type="callout-text"><a target="_blank" rel="noopener noreferrer nofollow" class="notion-link-token notion-focusable-token notion-enable-hover" href="http://go-colly.org/" style="pointer-events: none">Colly</a> is a Go package that allows you to write both web scrapers and crawlers. It is built on top of Go's net/HTTP package for network communication, and <code>goquery</code>, which provides a "jQuery-like" syntax for targeting HTML elements.</div>
</div>

### Scraper logic

We are ready to start. Create a `scraper.go` file in your directory, and set up your `main` function.

```go
package main

import (
    // import Colly
	"github.com/gocolly/colly"
)

func main(){
}
```

<div data-node-type="callout">
<div data-node-type="callout-emoji">🧠</div>
<div data-node-type="callout-text">If you haven't used Go before, this might look a bit weird. If you're asking yourself what is main? where does that package come from? then check out my other blog post that goes into detail on this.</div>
</div>

**Using Colly**

The bulk of this program will be handled by Colly and the Collector object which manages the network communication and is responsible for the execution of the attached **callbacks**.

To use Let's initialise a Collector inside your main function:

`collector := colly.NewCollector()`

You can customise the parameters of your Collector to suit different needs. We will be using the default setup, but if you're curious you can check out the Colly documentation for more.

Firstly, let's start with the most simple case to see if our set-up is working: we want to visit a URL and print in our console the URL we are visiting. This is the URL I want to use: [https://ricette.giallozafferano.it/Schiacciata-fiorentina.html](https://ricette.giallozafferano.it/Schiacciata-fiorentina.html)

<div data-node-type="callout">
<div data-node-type="callout-emoji">💡</div>
<div data-node-type="callout-text">The code in this program is catered to work on that URL (and other recipes from that website) - it won't work on any other URL so use the same URL if you are following along.</div>
</div>

**Using command-line arguments**

In every tutorial I checked out before building this scraper, they copy-pasted the url locally in the file and then passed it to the Collector calling `collector.Visit("their_url_string_here).`  
I didn't find that solution very elegant, and in the name of learning Go and experimenting, I looked up how to pass command line arguments to a Go app, so that I could pass the url as a command line argument, and scrape different urls from the same website without having to manually change the code.

To achieve this, we are using the `os` package, so we're going to add this to the import, and fill in our main function with what we need so far:

* A variable called `args` which will store the command line arguments we pass at run time.
    
* A variable called `url` to which we assign the value at index 1 from the `args` object. ❗️This is where we're storing the URL that we pass in the command line.
    
* A Colly collector object.
    

```go
package main

import (
    "os"
    // import Colly
	"github.com/gocolly/colly"
)

func main(){
    args := os.Args
	url := args[1]
	collector := colly.NewCollector()
}
```

<div data-node-type="callout">
<div data-node-type="callout-emoji">👀</div>
<div data-node-type="callout-text">For an in-depth example of how to work with command-line arguments in Go, you should check out this page from the "Go by example" site: <a target="_blank" rel="noopener noreferrer nofollow" href="" style="pointer-events: none">Go by Example: Command-Line Arguments</a>.</div>
</div>

Now that we have the building blocks of our program, let's try it out and see if it works by setting up callbacks for our collector object. Just so you know, everything we are going to write from now on is to be placed inside the `main` function. For our print statements, we are using the package `fmt` - it stands for "formatting" and you can import it in the import function as shown in the code block below.

```go
package main

import (
	"fmt"
	"os"
	// import Colly
	"github.com/gocolly/colly"
)
func main() {
	args := os.Args
	url := args[1]
	collector := colly.NewCollector()
	
  // whenever the collector is about to make a new request
	collector.OnRequest(func(r *colly.Request) {
        // print the url of that request
		fmt.Println("Visiting", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Blimey, an error occurred!:", e)
	})
	collector.Visit(url)

}
```

<div data-node-type="callout">
<div data-node-type="callout-emoji">🔎</div>
<div data-node-type="callout-text">If you haven't used callbacks before, here's a brief explanation of what is happening there. I'll use the first callback as an example (<code>OnRequest</code>), but the explanation applies to each of them.</div>
</div>

1. The `OnRequest` is a method of the Collector object, and it registers a function to be called whenever a new request is about to be made by the collector.
    
2. We are passing it an anonymous function that takes a `*colly.Request` object as an argument (representing the request that is about to be made). Inside this function, we are using the `fmt.Println` function to print out the message "Visiting" along with the URL of the request that is currently being processed in the command line.
    

Essentially, whenever the collector is about to **make a request**, this function will be executed, and it will print a message indicating that the collector is visiting a particular URL. This can be useful for debugging and monitoring the progress of the web scraping process, allowing you to see which URLs are being visited by the collector. Similarly, the `OnResponse` and `OnError` methods will trigger when the collector is about to receive a response or raise an error respectively.

In your terminal, run

```bash
go run scraper.go https://ricette.giallozafferano.it/Schiacciata-fiorentina.html
```

and you should see the following being printed back:

```bash
Visiting https://ricette.giallozafferano.it/Schiacciata-fiorentina.html
Got a response from https://ricette.giallozafferano.it/Schiacciata-fiorentina.html
```

**🎨 Now comes the fun**

To write our main logic, we are going to use the `OnHTML` callback. To fully understand this, I recommend you spend some quality time with the Colly documentation. The function that we define in this callback will be run when the collector encounters an HTML element that you can specify as a **css-selector**. In the example below, the function will be run when the collector finds an element with class name "div-main".

```go
collector.OnHTML(".div-main", func(e *colly.HTMLElement) {
    // Your code to handle the HTML element matching the selector
})
```

To know which css-selector to target, we'll need to spend some time inspecting the webpage we are going to scrape. To do this, use your browser web inspector and look at the elements of the content you want to extract.

While you can multiple OnHtml callbacks in your program, each handling a different css-selector for the multiple elements you want to scrape, I felt it was cleaner to have one callback that triggers for the `main` element of the page, and to handle the child elements scraping cascading within the same function.

Let's do a little recap before we move on:

* We have our collector
    
* We know how to scrape the url and we know how to pass the url
    

But what are we going to do with the data scraped?  
Since we are scraping a recipe, it only makes sense to create a `struct` to store the recipe features we are interested in extracting.

### **Defining Data Structures**

I spent some time looking at the recipe page and deciding which elements I wanted to extract: the recipe name, the recipe specs (such as difficulty, preparation time, cooking time, serving size and price tier), and the ingredients.

Firstly, I defined a `struct` for the recipe, with `url` and `name` as strings. Then I had to think of which data structure to use for the **recipe specs** and the **ingredients**. Different recipes will have different combinations of ingredients, but will always have the same types of specs, just with different values, so it made sense to use different data structures for the two.

* As the recipe specs are **fixed** fields across different recipes, such as the name and URL fields, I decided to create a separate struct to hold the `RecipeSpecs`
    
* For the ingredients, however, I needed a more **flexible** data structure to which I could append key-value pairs, as I didn't know how many or which keys I would encounter in the ingredient list. To achieve this, I created a Dictionary object of type `map` which defines a mapping of key-value pairs from string to string type.
    

```go
type Dictionary map[string]string

type RecipeSpecs struct {
	difficulty, prepTime, cookingTime, servingSize, priceTier string
}

type Recipe struct {
	url, name      string
	ingredients    []Dictionary
	specifications RecipeSpecs
}
```

Phew, with that out of the way, I was ready to save the scraped data to my struct.  
Have a look at the code below and in-line comments for explanations of how I put it together. Note that the recipe is in Italian, but thanks to the Latin influence onto the English language, a lot of the words are very similar to English:

* Difficoltà: difficulty
    
* Preparazione: preparation
    
* Cottura: cooking (aka cooking time)
    
* Dosi per: doses for (aka serving size)
    
* Costo: cost
    

```go
// initialise a slice of type Recipe (like a list of recipes)
// this way we'll be able to append each recipe to it, 
// and access the recipes outside the scope of the callback function.
var recipes []Recipe 

  
c.OnHTML("main", func(main *colly.HTMLElement) {
        // initialise a new recipe struct every time we visit a page
		recipe := Recipe{}
        // initialise a new Dictionary object to stoer the ingredients mappings
		ingredients_dictionary := Dictionary{}
    
        // assign the value of URL (the url we are visiting) to the recipe field
		recipe.url = url

		// find the recipe title, assign it to the struct, and print it in the command line
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

// finally, run the scraper 
collector.Visit(url)
```

## Conclusion 🏁

And this is it! Congratulations! You've successfully built a web scraper in Go using the Colly library, and learned some Italian words 🇮🇹  
You can now apply this knowledge to scrape data from other websites or enhance the scraper's capabilities.

Remember that web scraping should be done responsibly and in accordance with the website's terms of use. Always be respectful of the website's resources and consider implementing rate limiting and error handling mechanisms to ensure a smooth scraping process if you are building a more complex application.

I will revisit this project in the future and look to implement more functions, such as:

* Add a module to save the scraped data locally (such as a csv file).
    
* Cache pages so subsequent runs don't have to download the same page again.
    

What else should I implement? And did you enjoy this blog post? Leave a comment with your thoughts, and see you next time 👋

<div data-node-type="callout">
<div data-node-type="callout-emoji">💙</div>
<div data-node-type="callout-text"><strong>You can find the full codebase on my Github in the Webscraper repository.</strong></div>
</div>
