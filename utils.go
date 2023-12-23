package main

import (
	"log"
	"regexp"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/mmcdole/gofeed"
)

const regex = `<.*?>`

// This method uses a regular expression to remove HTML tags.
func stripHtmlRegex(s string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(s, "")
}

// This method uses a regular expression to remove HTML tags.
func htmlToMarkdown(s string) string {
	converter := md.NewConverter("", true, nil)
	text, err := converter.ConvertString(s)
	if err != nil {
		log.Fatal(err)
	}
	return text
}

func reverseArray(arr []*gofeed.Item) {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func truncateString(s string, maxLen int) string {
    runes := []rune(s)
    if len(runes) <= maxLen {
        return s
    }
    if maxLen < 3 {
        maxLen = 3
    }
    return string(runes[0:maxLen-1]) + "…"
}
