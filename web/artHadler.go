package handler

import (
	"ascii-art-web/art"
	"html/template"
	"net/http"
	"strings"
)

type Output struct {
	Pic string
}

func ArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, r, errStatus{http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)})
		return
	}
	t, err := template.ParseFiles("./static/tmp/index.html")
	if err != nil {
		ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}
	err = r.ParseForm()
	if err != nil {
		ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}
	message := r.PostFormValue("input")
	font := r.PostFormValue("font")
	font += ".txt"
	message = strings.ReplaceAll(message, "\r", "")
	message, status := art.Ascii(message, font)
	if status != 0 {
		switch status {
		case http.StatusBadRequest:
			ErrorHandler(w, r, errStatus{http.StatusBadRequest, http.StatusText(http.StatusBadRequest)})
		default:
			ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		}
		return
	}
	art := Output{message}
	err = t.ExecuteTemplate(w, "index", art)
	if err != nil {
		ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}
}
