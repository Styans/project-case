package handlers

import (
	"ascii-art/internal/drawtext"
	"html/template"
	"net/http"
)

type FormDatas struct {
	input string
	fs    string
}

func (app *Aplication) mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.errors(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tmlp, err := template.ParseFiles("./internal/web/html/index.html", "./internal/web/html/header.html")
		if err != nil {
			app.errors(w, http.StatusNotFound)
			return
		}
		err = tmlp.ExecuteTemplate(w, "index", nil)
		if err != nil {
			app.errors(w, http.StatusNotFound)
			return
		}
	case http.MethodPost:
		r.ParseForm()

		res := drawtext.TextConv(r.FormValue("input"), r.FormValue("fs"), app.Fonts)

		tmlp, err := template.ParseFiles("./internal/web/html/index.html", "./internal/web/html/header.html")
		if err != nil {
			app.errors(w, http.StatusNotFound)
			return
		}
		err = tmlp.ExecuteTemplate(w, "index", res)
		if err != nil {
			app.errors(w, http.StatusNotFound)
			return
		}
	default:
		w.Header().Set("Allow", http.MethodPost)
		app.errors(w, http.StatusMethodNotAllowed)
		return
	}

}
