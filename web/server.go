package handler

import (
	"log"
	"net/http"
)

func Handler() {
	log.Printf("Tap to link http://localhost:3000")

	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./static/style"))))
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/ascii-art", ArtHandler)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
