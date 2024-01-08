package TPBlog

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strings"
)

var filename = path + "content/articles.json"

func RetrieveArticles() ([]Article, error) {
	var articles []Article

	data, err := os.ReadFile(filename)

	if len(data) == 0 {
		return nil, nil
	}

	err = json.Unmarshal(data, &articles)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func searchArticle(search string) []Article {
	var result []Article

	articles, err := RetrieveArticles()
	if err != nil {
		log.Fatal("log: RetrieveArticles() error!\n", err)
	}

	for _, article := range articles {
		if strings.Contains(article.Title, search) {
			result = append(result, article)
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result
}

func changeArticles(articles []Article) {
	data, errJSON := json.Marshal(articles)
	if errJSON != nil {
		log.Fatal("log: addArticle()\t JSON Marshall error!\n", errJSON)
	}
	errWrite := os.WriteFile(filename, data, 0666)
	if errWrite != nil {
		log.Fatal("log: addArticle()\t WriteFile error!\n", errWrite)
	}
}

func addArticle(newCtn Article) {
	// Don't forget to add automatically the article id in the addArticle Handler!
	articles, err := RetrieveArticles()
	if err != nil {
		log.Fatal("log: RetrieveArticles() error!\n", err)
	}
	articles = append(articles, newCtn)
	changeArticles(articles)
}

func deleteArticle(id int) {
	articles, err := RetrieveArticles()
	if err != nil {
		log.Fatal("log: RetrieveArticles() error!\n", err)
	}
	for i, article := range articles {
		if article.Id == id {
			articles = append(articles[:i], articles[i+1:]...)
		}
	}
	changeArticles(articles)
}

func selectCategory(category string) []Article {
	var selectArticles []Article
	articles, err := RetrieveArticles()
	if err != nil {
		log.Fatal("log: RetrieveArticles() error!\n", err)
	}
	for _, article := range articles {
		if article.Category == category {
			selectArticles = append(selectArticles, article)
		}
	}
	return selectArticles
}

func randomArticles() []Article {
	articles, err := RetrieveArticles()
	if err != nil {
		log.Fatal("log: RetrieveArticles() error!\n", err)
	}
	if len(articles) < 10 {
		return articles
	}
	rand.Shuffle(len(articles), func(i, j int) {
		articles[i], articles[j] = articles[j], articles[i]
	})
	return articles[:10]
}

func modifyArticle(updatedArticle Article) {
	articles, err := RetrieveArticles()
	if err != nil {
		log.Fatal("log: RetrieveArticles() error!\n", err)
	}
	for i, article := range articles {
		if article.Id == updatedArticle.Id {
			articles[i] = updatedArticle
		}
	}
	changeArticles(articles)
}
