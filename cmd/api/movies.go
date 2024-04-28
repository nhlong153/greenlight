package main

import (
	"fmt"
	"net/http"
)

type Movie struct {
	Title   string   `json:"title"`
	Year    int32    `json:"year"`
	Runtime int32    `json:runtime`
	genres  []string `json:"genres`
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	movie := Movie{}
	err := app.readJson(w, r, &movie)
	if err != nil {
		return
	}
	fmt.Fprintln(w, "Create a movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Show movie with id %d\n", id)

}
