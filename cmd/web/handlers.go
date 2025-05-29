package main

import (
	"log"
	"net/http"

	ui "github.com/patvoj/snippetbox/ui/html/pages"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := ui.Home().Render(r.Context(), w)
	if err != nil {
		log.Printf("error rendering home: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
