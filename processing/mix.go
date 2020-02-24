package processing

import (
	"github.com/NatalyMac/test/model"
)

// MixArticles mixes regular and content marketing articles
func MixArticles(articles, cmArticles []model.Article) []model.Article {
	n := 0
	m := 0
	var items, aSlice, cmSlice []model.Article
	for n < len(articles) {
		aSlice = articles[n : n+5]
		if m <= len(cmArticles)-1 {
			cmSlice = cmArticles[m : m+1]
		} else {
			ad := model.Article{
				Type: "Ad",
			}
			cmSlice = make([]model.Article, 1)
			cmSlice[0] = ad
		}
		items = append(items, aSlice...)
		items = append(items, cmSlice...)
		n += 5
		m++
		if n > len(articles) {
			length := len(items) - (n - len(articles))
			items = items[:length-1]
			break
		}
	}
	return items
}
