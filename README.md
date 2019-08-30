# Golang

##Given a string substr, getMovieTitles must perform the following tasks:

**Query - https://jsonmock.hackerrank.com/api/movies/search/?Title=substr(replace substr)**

Initialize the titles array to store total string elements. Store the Title of each movie
meeting the search criterion in the titles array.
Sort titles in ascending order and return it as your answer .
The query response from the website is a JSON response with the following five
fields:

page : The current page.
per_page : The maximum number of results per page.
total : The total number of movies in the search result.
total_pages : The total number of pages which must be queried to get all the results.
data : An array of JSON objects containing movie information where the Title field
denotes the title of the movie.

In order to get all results, you may have to make multiple page requests. To request a
page by number, your query should read
https://jsonmock.hackerrank.com/api/movies/search/?Title=substr&page=pageNumber 
, replacing substr and pageNumber.

Expacted Output:
1. When we are passing "movie name", "page number" we should get movie title
array with sorted order(If we are not passing page then by default it should consider
1)
2. Write a function which will accept "movie name", "page number", "Year" and will
return the movie title array with sorted order(If we are not passing page then by
default it should consider 1)