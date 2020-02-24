package api

import (
	"encoding/json"
	"net/http"

	"github.com/NatalyMac/test/model"
	"github.com/NatalyMac/test/fetch"
	"github.com/NatalyMac/test/processing"
)

// GetMix gets regular and cm articles, calls mixer and returns mixed
func GetMix(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	articles, cmArticles := fetch.GetContent()
	items := processing.MixArticles(articles, cmArticles)
	var res model.Response
	res.HTTPStatus = 200
	res.Response.Items = items
	json.NewEncoder(w).Encode(res)
}
