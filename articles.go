package TPBlog

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

var jsonfile = path + "content/articles.json"

//	retrieveArticles
//
// retrieves all Article present in articles.json and stores them in a slice of Article.
// It returns the slice of Article and an error.
func retrieveArticles() ([]Article, error) {
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

//	searchArticles
//
// retrieves all Article in which Article.Title contains `search` and returns them in a slice.
func searchArticle(search string) []Article {
	var result []Article

	articles, err := retrieveArticles()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
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

//	changeArticles
//
// overwrites articles.json with `articles` in json format.
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

//	getIdNewArticle
//
// returns first unused id in articles.json.
func getIdNewArticle() int {
	articles, err := retrieveArticles()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	var id int
	var idFound bool
	for id = 1; !idFound; id++ {
		idFound = true
		for _, article := range articles {
			if article.Id == id {
				idFound = false
			}
		}
	}
	id--
	return id
}

//	addArticle
//
// adds the Article `newCtn` to articles.json.
func addArticle(newCtn Article) {
	articles, err := retrieveArticles()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	articles = append(articles, newCtn)
	changeArticles(articles)
}

//	deleteArticle
//
// remove the Article which Article.Id is sent in argument from articles.json.
func deleteArticle(id int) {
	articles, err := retrieveArticles()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	for i, article := range articles {
		if article.Id == id {
			articles = append(articles[:i], articles[i+1:]...)
		}
	}
	changeArticles(articles)
}

//	selectCategory
//
// returns all Article which Article.Category matches the `category` argument.
func selectCategory(category string) []Article {
	var selectArticles []Article
	articles, err := retrieveArticles()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	for _, article := range articles {
		if article.Category == category {
			selectArticles = append(selectArticles, article)
		}
	}
	return selectArticles
}

// selectArticle
// returns the Article which Article.Id matches the `id` argument.
func selectArticle(id int) (Article, bool) {
	var article Article
	articles, err := retrieveArticles()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	var ok bool
	for _, singleArticle := range articles {
		if singleArticle.Id == id {
			ok = true
			article = singleArticle
		}
	}
	return article, ok
}

//	randomArticles
//
// select randomly a fixed number of Article and returns them in a slice.
func randomArticles() []Article {
	articles, err := retrieveArticles()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	if len(articles) < 12 {
		return articles
	}
	rand.Shuffle(len(articles), func(i, j int) {
		articles[i], articles[j] = articles[j], articles[i]
	})
	return articles[:12]
}

//	modifyArticle
//
// modifies the Article in articles.json that matches
// `updatedArticle`'s Id with `updatedArticle`'s content.
func modifyArticle(updatedArticle Article) {
	articles, err := retrieveArticles()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	for i, article := range articles {
		if article.Id == updatedArticle.Id {
			articles[i] = updatedArticle
		}
	}
	changeArticles(articles)
}

//	formatArticle
//
// replace all # markdown title signs with html in `article`'s content and returns the Article 's formatted version.
func formatArticle(article Article) ArticleHTML {
	var Match bool = true
	var formattedArticle ArticleHTML
	title3 := regexp.MustCompile("### .+\\n")
	titles3 := title3.FindAllString(article.Content, -1)
	if len(titles3) == 0 {
		Match = false
	}
	if Match {
		for i, title := range titles3 {
			title = strings.Replace(title, "### ", "<div class=\"ctn-title3\">", 1)
			title = strings.Replace(title, "\n", "</div>\n", 1)
			article.Content = strings.Replace(article.Content, titles3[i], title, 1)
		}
	}
	Match = true
	title2 := regexp.MustCompile("## .+\\n")
	titles2 := title2.FindAllString(article.Content, -1)
	if len(titles2) == 0 {
		Match = false
	}
	if Match {
		for i, title := range titles2 {
			title = strings.Replace(title, "## ", "<div class=\"ctn-title2\">", 1)
			title = strings.Replace(title, "\n", "</div>\n", 1)
			article.Content = strings.Replace(article.Content, titles2[i], title, 1)
		}
	}
	Match = true
	title1 := regexp.MustCompile("# .+\\n")
	titles1 := title1.FindAllString(article.Content, -1)
	if len(titles1) == 0 {
		Match = false
	}
	if Match {
		for i, title := range titles1 {
			title = strings.Replace(title, "# ", "<div class=\"ctn-title1\">", 1)
			title = strings.Replace(title, "\n", "</div>\n", 1)
			article.Content = strings.Replace(article.Content, titles1[i], title, 1)
		}
	}

	formattedArticle = ArticleHTML{
		Id:           article.Id,
		Category:     article.Category,
		Title:        article.Title,
		Author:       article.Author,
		Date:         article.Date,
		BigImg:       article.BigImg,
		SmallImg:     article.SmallImg,
		Introduction: article.Introduction,
		Content:      template.HTML(article.Content),
	}

	return formattedArticle
}
