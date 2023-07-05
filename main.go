package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
    "github.com/spf13/viper"
)



type Movies struct {
	Search []struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	} `json:"Search"`
	TotalResults string `json:"totalResults"`
	Response     string `json:"Response"`
}
type Movie  struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	Dvd        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}

func main(){
    viper.SetConfigFile(".env")
    viper.ReadInConfig()
    api_key := fmt.Sprint(viper.Get("API_KEY"))
    args := os.Args[1:]
    if len(args) == 0 {
        log.Fatal("No aruguments supplied")
    }
    var query string 
    for i, value := range(args){
        if i == 0{
            query += fmt.Sprint(value)
            continue
        }
        query += fmt.Sprint(" ",value) 
    }
    uniqueID := firstQuery(query, api_key)
    movieDetails := finalQuery(uniqueID, api_key)
    printDetails(movieDetails)
}

func firstQuery(movieName, key string) string{
    queryString := fmt.Sprintf("http://www.omdbapi.com/?apikey="+key+"&s="+movieName+"&plot=full")
    res, err := http.Get(queryString)
    if err != nil {
        log.Fatalln(err)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatalln(err)
    }
    var results Movies 
    if err := json.Unmarshal(body, &results); err!= nil{
        log.Fatalln("Failed to unmarshal data on first query")
    }
    if results.Response == "False"{
        log.Fatalln("The movie could not be fetched. Response failure at first query") 
    }    
    return results.Search[0].ImdbID
}
func finalQuery(movieID,key string) Movie{
    queryString := fmt.Sprintf("http://www.omdbapi.com/?apikey="+key+"&i="+movieID+"&plot=full")
    res, err := http.Get(queryString)
    if err != nil {
        log.Fatalln(err)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatalln(err)
    }
    var result Movie 
    if err := json.Unmarshal(body, &result); err!= nil{
        log.Fatalln("Failed to unmarshal data on first query")
    }
    return result
}



// This function takes in a Movie Struct and prints it 
func printDetails(data Movie){
    fmt.Println(data.Title)
    fmt.Println("Directed by:", data.Director)
    fmt.Println("Runtime:",data.Runtime)
    fmt.Println("Actors:", data.Actors)
    fmt.Println("IMDB Rating",data.ImdbRating)
    fmt.Println("Metascore", data.Metascore)
    fmt.Println("plot:",data.Plot)
    fmt.Println("Poster:", data.Poster)
    
}

