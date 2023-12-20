package main

import (
	"log"
	"regexp"
	md "github.com/JohannesKaufmann/html-to-markdown"
)

const regex = `<.*?>`

// This method uses a regular expresion to remove HTML tags.
func stripHtmlRegex(s string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(s, "")
}

// This method uses a regular expresion to remove HTML tags.
func htmlToMarkdown(s string) string {
	converter := md.NewConverter("", true, nil)
	text, err := converter.ConvertString(s)
	if err != nil {
		log.Fatal(err)
	}
	return text
}
