package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
	alphabet      map[rune]string
	ihfa          map[rune]bool
	qamariya      map[rune]bool
	hards         map[rune]bool
	dataSlice     []*templateData
}

func main() {
	// addr := flag.String("addr", ":4000", "Сетевой адрес веб-сервера")
	port := os.Getenv("PORT")
	port = ":" + port
	addr := flag.String("addr", port, "Сетевой адрес веб-сервера")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
	}
	app.alphabet = app.initAlphabet()
	app.ihfa = app.initIhfa()
	app.qamariya = app.initQamariya()
	app.hards = app.initHards()
	app.dataSlice = app.initDataSlice()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запуск сервера на http://127.0.0.1%s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
