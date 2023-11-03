/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import (
	"encoding/json"
	"strings"
)

func GetResultsJsonString(totalResults float64, channelIds map[string]bool, query string) (string, error) {
	channelIdsArray := make([]string, len(channelIds))
	i := 0
	for k := range channelIds {
		channelIdsArray[i] = k
		i++
	}
	totalResultsInt := int(totalResults)
	channelIdsString := strings.Join(channelIdsArray, ",")
	result := map[string]any{
		"totalResults": totalResultsInt,
		"channelIds":   channelIdsString,
		"searchTerm":   query,
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultJson), nil
}
