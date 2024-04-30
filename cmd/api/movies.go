package main

import (
	"fmt"
	"net/http"

	"github.com/nhlong153/greenlight/internal/data"
)

type Movie struct {
	Title   string   `json:"title"`
	Year    int32    `json:"year"`
	Runtime int32    `json:runtime`
	genres  []string `json:"genres`
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}
	err := app.readJson(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "Error Parsing data")
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: int64(input.Runtime),
		Genres:  input.Genres,
	}

	err = app.models.Movies.Insert(movie)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, movie, nil)
	if err != nil {

		app.serverErrorResponse(w, r, err)

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

	movie, err := app.models.Movies.Get(id)
	if err != nil {
		app.logError(r, err)
		app.notFoundResponse(w, r)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
