package maincrudsql

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Title    string    `json:"title"`
	Rating   int       `json:"rating"`
	Link     string    `json:"link"`
	Director *Director `json:"director"`
	Genres   string    `json:"genre"`
}
type Director struct {
	FirstName string
	LastName  string
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Title == params["title"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
func creatMovie() {

}
func updateMovie() {

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{Title: "new", Rating: 20, Link: "Fake", Director: &Director{FirstName: "Zad", LastName: "Amumum"}, Genres: "Sad"})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{title}", getMovies).Methods("GET")
	r.HandleFunc("/movies", creatMovie).Methods("POST")
	r.HandleFunc("/movies/{title}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{title}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at Port 80000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
