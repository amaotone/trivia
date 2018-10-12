package main

import (
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestParseDocument(t *testing.T) {
	cases := []string{
		"testdata/wikipedia_en.html",
		"testdata/wikipedia_ja.html",
	}
	for _, c := range cases {
		fp, _ := os.Open(c)
		defer fp.Close()
		doc, _ := goquery.NewDocumentFromReader(fp)
		title, lead := parseDocument(doc)
		if title == "" || lead == "" {
			t.Error("title or lead is empty", title, lead)
		}
	}
}
