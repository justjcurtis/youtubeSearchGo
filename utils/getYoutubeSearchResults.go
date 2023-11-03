/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

const URL string = "https://www.googleapis.com/youtube/v3/search?maxResults=20&part=snippet&q="

func mergeResults(a map[string]bool, b map[string]bool) map[string]bool {
	for k := range b {
		a[k] = true
	}
	return a
}

func GetYoutubeSearchResults(query string, pageToken string) (float64, map[string]bool, error) {
	key, err := GetEnvKey("API_KEY")
	if err != nil {
		return 0, nil, err
	}

	url := URL + query + "&key=" + key
	if pageToken != "" {
		url += "&pageToken=" + pageToken
	}

	resp, err := http.Get(url)
	if err != nil {
		return 0, nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	data := make(map[string]any)
	json.Unmarshal(body, &data)

	items, ok := data["items"]
	if !ok {
		return 0, nil, err
	}

	itemsArray, ok := items.([]any)
	if !ok {
		return 0, nil, err
	}

	channelTitles := make(map[string]bool)
	for _, item := range itemsArray {
		snippet, ok := item.(map[string]any)["snippet"]
		if !ok {
			return 0, nil, err
		}
		channelTitle, ok := snippet.(map[string]any)["channelTitle"]
		if !ok {
			return 0, nil, err
		}
		channelTitles[channelTitle.(string)] = true
	}

	nextPageToken, ok := data["nextPageToken"]
	if ok && nextPageToken != "" {
		_, next, err := GetYoutubeSearchResults(query, nextPageToken.(string))
		if err != nil {
			return 0, nil, err
		}
		channelTitles = mergeResults(channelTitles, next)
	}

	if pageToken == "" {
		pageinfo, ok := data["pageInfo"]
		if !ok {
			return 0, nil, err
		}
		totalResults, ok := pageinfo.(map[string]any)["totalResults"]
		if !ok {
			return 0, nil, err
		}
		return totalResults.(float64), channelTitles, nil
	}

	return 0, channelTitles, nil
}
