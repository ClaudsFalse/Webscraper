# Webscraper
A "GialloZafferano" web scraper command line application built in Go using Colly.
The app is designed to scrape the recipe pages of the GialloZafferano website.
This is not affiliated with GialloZafferano, I just like italian food. 

# How to run
Run `go run main/scraper.go <url-of-the-recipe> ` 

# What's next?
What I plan to add to the package:
- Add a module to save the scraped data locally (such as a csv file).
- Cache pages so subsequent runs don't have to download the same page again.

