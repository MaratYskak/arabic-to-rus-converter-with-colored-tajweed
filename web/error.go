package handler

import (
	"html/template"
	"net/http"
)

type errStatus struct {
	StatusCode int
	StatusText string
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status errStatus) {
	t, err := template.ParseFiles("./static/tmp/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status.StatusCode)
	err = t.ExecuteTemplate(w, "error", status)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
