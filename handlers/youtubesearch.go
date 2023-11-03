/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package handlers

import (
	"net/http"
	"youtubesearch/utils"
)

func HandleYoutubeSearch(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	hasQuery := params.Has("q")
	if !hasQuery {
		return
	}
	query := params.Get("q")
	totalResults, channelTitles, err := utils.GetYoutubeSearchResults(query, "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	resultJson, err := utils.GetResultsJsonString(totalResults, channelTitles, query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resultJson))
}
