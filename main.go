package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

//Movie defines exported struct type of movie details
type Movie struct {
	Title  string `json:"Title"`
	Year   int64  `json:"Year"`
	ImdbID string `json:"imdbID"`
}

//Records defines exported struct type of movie records
type Records struct {
	Page       string  `json:"page"`
	PerPage    int64   `json:"per_page"`
	Total      int64   `json:"total"`
	TotalPages int64   `json:"total_pages"`
	Data       []Movie `json:"data"`
}

//DefaultPageNumber defines constant of type string
const DefaultPageNumber = "1"

func getMovieTitles(movieTitle string, year string, pageNumber string) {
	if pageNumber == "" {
		pageNumber = DefaultPageNumber
	}

	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/movies/search?Title=%s&Year=%s&page=%s", movieTitle, year, pageNumber)
	response, responseError := http.Get(url)
	if responseError != nil {
		fmt.Printf("The HTTP request failed with error %s", responseError)
	} else {
		responseBody, _ := ioutil.ReadAll(response.Body)

		var movieRecords = new(Records)
		parseError := json.Unmarshal(responseBody, &movieRecords)
		if parseError != nil {
			fmt.Println("Error occurs while parsing json response ", parseError)
		}

		titles := make([]string, len(movieRecords.Data))
		for index, value := range movieRecords.Data {
			titles[index] = value.Title
		}

		if len(titles) != 0 {
			sort.Strings(titles)
			fmt.Println("\n *** MOVIE LIST ***")
			for _, title := range titles {
				fmt.Println(" ", title)
			}
		} else {
			fmt.Println("\nNo records found for given search criteria")
		}
	}
}

func main() {
	var movieTitle string
	var page string
	var year string

	fmt.Print("Please Enter the Movie Title subString: ")
	fmt.Scanf("%s", &movieTitle)

	fmt.Print("Page Number: ")
	fmt.Scanf("%s", &page)

	fmt.Print("Year: ")
	fmt.Scanf("%s", &year)

	getMovieTitles(movieTitle, year, page)
}
