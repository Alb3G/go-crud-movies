package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"Isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	// Después de borrar la pelicula devolvemos las peliculas existentes sin la eliminada
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
}

func main() {
	r := mux.NewRouter()

	// a HandleFunc siempre le decimos primero la ruta en la que vamos a llamar a la funcion en este caso getMovies.

	movies = append(
		movies,
		Movie{
			ID:    "1",
			Isbn:  "438227",
			Title: "Movie One",
			Director: &Director{
				FirstName: "Jon",
				LastName:  "Doe",
			},
		},
	)

	movies = append(
		movies,
		Movie{
			ID:    "2",
			Isbn:  "45455",
			Title: "Movie Two",
			Director: &Director{
				FirstName: "Steve",
				LastName:  "Smith",
			},
		},
	)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
