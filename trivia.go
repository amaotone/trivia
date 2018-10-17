package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var (
	flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Usage: "This flag specify the language to search.\n\t" + "Available languages: https://meta.wikimedia.org/wiki/List_of_Wikipedias",
		},
	}
)

// os exit codes
const (
	ExitCodeOk = iota
	ExitCodeError
)

func parseConfig(c *cli.Context) (lang string) {
	config := loadConfig()
	lang = c.String("lang")
	if lang == "" {
		if config.Lang != "" {
			lang = config.Lang
		} else {
			lang = "en" // default language
		}
	}
	return lang
}

func parseDocument(doc *goquery.Document) (title string, lead string) {
	title = doc.Find("#firstHeading").Text()
	lead = doc.Find("#mw-content-text > div > p").First().Text()
	return title, lead
}

func fetchWord(lang string) (title string, lead string) {
	url := fmt.Sprintf("http://%s.wikipedia.org/wiki/Special:Randompage", lang)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("Wikipedia scraping failed.")
		os.Exit(ExitCodeError)
	}
	title, lead = parseDocument(doc)
	return title, lead
}

func showWord(c *cli.Context) {
	lang := parseConfig(c)
	title, lead := fetchWord(lang)
	bold := color.New(color.Bold)
	bold.Println(strings.TrimSpace(title))
	fmt.Println(strings.TrimSpace(lead))
	os.Exit(ExitCodeOk)
}

func main() {
	app := cli.NewApp()
	app.Name = "Trivia"
	app.Usage = "Trivia makes your life richer."
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:   "set",
			Usage:  "Save config to $HOME/.trivia/config.json",
			Action: setConfig,
			Flags:  flags,
		},
	}
	app.Action = showWord
	app.Flags = flags
	app.Run(os.Args)
}
