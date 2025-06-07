package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/patvoj/snippetbox/internal/models"
	ui "github.com/patvoj/snippetbox/ui/html"
	pages "github.com/patvoj/snippetbox/ui/html/pages"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, 200, pages.Home(data))
}

func (app *application) getSnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.serverError(w, r, err)
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}

		return
	}

	snippet.Content = strings.ReplaceAll(snippet.Content, "\\n", "\n")
	title := "Snippet #" + strconv.Itoa(snippet.ID)

	data := app.newTemplateData(r)
	data.Snippet = &snippet

	snippetViewComponent := pages.SnippetView(data)

	app.render(w, r, 200, ui.Base(title, data.CurrentYear, snippetViewComponent))
}

func (app *application) getSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a snippet..."))
}

func (app *application) postSnippetCreate(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
