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
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Name string `json:"name"`
	Director *Director `json:"director"`
} 

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var Movies[] Movie 

func appendMovie(id, isbn, name, directorFirst, directorLast string) {
	director := &Director{
		Firstname: directorFirst,
		Lastname:  directorLast,
	}

	movie := Movie{
		ID:        id,
		Isbn:      isbn,
		Name:      name,
		Director: director,
	}

	
	Movies = append(Movies, movie)
}
 

	func getMovies(w http.ResponseWriter , r *http.Request ){
		w.Header().Set("content-type","application/json")
		json.NewEncoder(w).Encode(Movies)
	}

	func deleteMovie(w http.ResponseWriter , r *http.Request){
		w.Header().Set("Content-Type","application/json")
		params := mux.Vars(r)
		for index , item := range Movies{
			if item.ID == params["id"]{
				Movies = append(Movies[:index],Movies[index+1:]... )
				break
			}
		}
		json.NewEncoder(w).Encode(Movies)
	}

	func getMovie(w http.ResponseWriter , r *http.Request){
		w.Header().Set("Content-Type","application/json")
		params := mux.Vars(r)
		for _,item := range Movies {
			if item.ID == params["id"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	}

	func createMovie(w http.ResponseWriter , r *http.Request){
		w.Header().Set("Content-Type" , "application/json")
		var movie Movie
		json.NewDecoder(r.Body).Decode(&movie)
		movie.ID = strconv.Itoa(rand.Intn(1000000))
		
		Movies = append(Movies, movie)

		json.NewEncoder(w).Encode(Movies)
	}

	func updateMovie(w http.ResponseWriter , r *http.Request){
		w.Header().Set("Content-Type","application/json")
		params := mux.Vars(r)
		var movie Movie
		json.NewDecoder(r.Body).Decode(&movie)

		for index , item := range Movies {
			if item.ID == params["id"] {
				Movies[index] = movie
				break
			}

		}
		json.NewEncoder(w).Encode(Movies)
	
	}



func main (){
	r := mux.NewRouter()

	appendMovie("1", "978-3-16-148410-0", "Inception", "Christopher", "Nolan")
	appendMovie("2", "978-1-23-456789-0", "The Matrix", "Lana", "Wachowski")
	appendMovie("3", "978-0-12-345678-9", "Interstellar", "Christopher", "Nolan")
	appendMovie("4", "978-0-98-765432-1", "The Dark Knight", "Christopher", "Nolan")
	appendMovie("5", "978-0-54-567890-1", "The Godfather", "Francis", "Coppola")
	appendMovie("6", "978-1-23-987654-0", "Pulp Fiction", "Quentin", "Tarantino")
	appendMovie("7", "978-1-23-654321-0", "Fight Club", "David", "Fincher")

	r.HandleFunc("/movies", getMovies ).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie/{id}",deleteMovie).Methods("DELETE")
	r.HandleFunc("/movie/{id}",updateMovie).Methods("PUT")

	fmt.Println("server is starting ....")
	log.Fatal(http.ListenAndServe(":8000",r))

}