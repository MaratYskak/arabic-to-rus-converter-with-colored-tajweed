package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Используем помощника render() для отображения шаблона.
	app.render(w, r, "home.page.html", nil)
}

func (app *application) howto(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/howto" {
		app.notFound(w)
		return
	}

	// Используем помощника render() для отображения шаблона.
	app.render(w, r, "howto.page.html", nil)
}

func (app *application) result(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/result" {
		app.notFound(w)
		return
	}

	text := r.FormValue("input")

	ArabicText := []rune(text)

	app.convert(ArabicText)

	// Используем помощника render() для отображения шаблона.
	app.render(w, r, "convert.page.html", app.dataSlice)
}
