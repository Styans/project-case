package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Aplication struct {
	Fonts map[string][]string
}

func (app *Aplication) Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.mainPage)
	style := http.FileServer(http.Dir("./internal/web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", style))
	return mux
}

func (app *Aplication) errors(w http.ResponseWriter, problem int) {
	e := "problem is " + strconv.Itoa(problem)
	w.WriteHeader(problem)

	tmpl, err := template.ParseFiles("./internal/web/html/errors.html")

	if err != nil {
		fmt.Fprint(w, e)
		return
	}
	tmpl.Execute(w, e)
}
