package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	// gResult maps to the result document received from the search

	gResult struct {
		GsearchResultClass string `json:"gsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNotFormatting string `json:"titleNotFormatting"`
		Content            string `json:"content"`
	}

	//gResponse contains the top level document
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results`
		} `json:"responseData"`
	}
)

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	// issue the search against google
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	defer resp.Body.Close()

	// Decode the json responsse into our struct type
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println(gr)
}
