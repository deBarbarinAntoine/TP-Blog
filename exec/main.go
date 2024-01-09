package main

import (
	"TPBlog"
	"fmt"
	"log"
)

func main() {
	//TPBlog.Run()

	//Test RetrieveArticles()
	articles, err := TPBlog.RetrieveArticles()
	if err != nil {
		log.Fatal("log: RetrieveArticles() error!\n", err)
	}
	fmt.Printf("articles: %#v\n", articles)
}
