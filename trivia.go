package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func initApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Trivia"
	app.Usage = "Make your life a little more happy."
	app.Version = "0.0.1"
	app.Action = action
	return app
}

func action(c *cli.Context) {
	lang := "ja"
	url := fmt.Sprintf("http://%s.wikipedia.org/wiki/Special:Randompage", lang)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("wikipedia scraping failed.")
	}

	title := doc.Find("#firstHeading").Text()
	lead := doc.Find("#mw-content-text>div>p").First().Text()

	bold := color.New(color.Bold)
	bold.Println(strings.TrimSpace(title))
	fmt.Println(strings.TrimSpace(lead))
}

func main() {
	app := initApp()
	app.Run(os.Args)
}
