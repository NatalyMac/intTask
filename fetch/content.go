package fetch

import (
	"encoding/json"
	"log"

	"github.com/NatalyMac/test/config"
	"github.com/NatalyMac/test/model"
)

func fetchData(url string) model.Response {
	var a model.Response
	result, err := Get(url)
	if err != nil {
		log.Printf("Error fetch Article: %v", err)
		return a
	}
	err = json.Unmarshal(result, &a)
	if err != nil {
		log.Printf("Error unmarshal Article: %+v", err)
	}
	return a
}

// GetContent requests content from given urls and returns 2 arrays of articles regular and content marketing
func GetContent() ([]model.Article, []model.Article) {
	conf := config.GetConfig()
	urlA := conf.ArticlesURL
	urlCm := conf.CMArticlesURL
	aChannel := make(chan model.Response, 1)
	cmChannel := make(chan model.Response, 1)
	go func() {
		aChannel <- fetchData(urlA)
	}()
	go func() {
		cmChannel <- fetchData(urlCm)
	}()
	var articles, cmArticles model.Response
	articles = <-aChannel
	cmArticles = <-cmChannel
	return articles.Response.Items, cmArticles.Response.Items
}