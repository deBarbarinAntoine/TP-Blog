package TPBlog

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

var jsonfile = path + "content/articles.json"

func RetrieveArticles() ([]Article, error) {
	var articles []Article

	data, err := os.ReadFile(jsonfile)

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
		if strings.Contains(strings.ToLower(article.Title), strings.ToLower(search)) {
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
	errWrite := os.WriteFile(jsonfile, data, 0666)
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

func selectArticle(id int) Article {
	var article Article
	articles, err := RetrieveArticles()
	if err != nil {
		log.Fatal("log: RetrieveArticles() error!\n", err)
	}
	for _, singleArticle := range articles {
		if singleArticle.Id == id {
			article = singleArticle
		}
	}
	return article
}

func randomArticles() []Article {
	articles, err := RetrieveArticles()
	if err != nil {
		log.Fatal("log: RetrieveArticles() error!\n", err)
	}
	if len(articles) < 12 {
		return articles
	}
	rand.Shuffle(len(articles), func(i, j int) {
		articles[i], articles[j] = articles[j], articles[i]
	})
	return articles[:12]
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

func formatArticle(article Article) string {
	var noMatch bool
	var ctn string
	title3 := regexp.MustCompile("### .+\\n")
	titles3 := title3.FindAllString(article.Content, -1)
	if len(titles3) == 0 {
		noMatch = true
	}
	if !noMatch {
		noMatch = false
		for i, title := range titles3 {
			openMark := regexp.MustCompile("### ")
			endMark := regexp.MustCompile("\\n")
			title = openMark.ReplaceAllString(title, "<div class=\"ctn-title3\">")
			title = endMark.ReplaceAllString(title, "</div>\n")
			ctn = strings.Replace(article.Content, titles3[i], title, -1)
		}
	}
	title2 := regexp.MustCompile("## .+\\n")
	titles2 := title2.FindAllString(article.Content, -1)
	if len(titles2) == 0 {
		noMatch = true
	}
	if !noMatch {
		noMatch = false
		for i, title := range titles2 {
			openMark := regexp.MustCompile("## ")
			endMark := regexp.MustCompile("\\n")
			title = openMark.ReplaceAllString(title, "<div class=\"ctn-title2\">")
			title = endMark.ReplaceAllString(title, "</div>\n")
			ctn = strings.Replace(article.Content, titles2[i], title, -1)
		}
	}
	title1 := regexp.MustCompile("# .+\\n")
	titles1 := title1.FindAllString(article.Content, -1)
	if len(titles1) == 0 {
		noMatch = true
	}
	if !noMatch {
		noMatch = false
		for i, title := range titles1 {
			openMark := regexp.MustCompile("# ")
			endMark := regexp.MustCompile("\\n")
			title = openMark.ReplaceAllString(title, "<div class=\"ctn-title1\">")
			title = endMark.ReplaceAllString(title, "</div>\n")
			ctn = strings.Replace(article.Content, titles1[i], title, -1)
		}
	}
	return ctn
}
