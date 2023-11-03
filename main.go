/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package main

import (
	"net/http"
	"youtubesearch/handlers"
)

func main() {
	http.HandleFunc("/youtubesearch", handlers.HandleYoutubeSearch)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
