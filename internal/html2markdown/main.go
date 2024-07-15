package main

import (
	"fmt"
	"log"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	readability "github.com/go-shiori/go-readability"
)

func main() {
	url := "https://www.ruanyifeng.com/blog/2024/07/weekly-issue-307.html"

	article, err := readability.FromURL(url, 30*time.Second)

	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(article.Content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title : %s\n", article.Title)
	fmt.Println("Markdwon : %s\n", markdown)
}
