package main

import (
	"bytes"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/a-h/templ"

	"github.com/patvoj/snippetbox/internal/utils"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, component templ.Component) {
	buf := new(bytes.Buffer)

	err := component.Render(r.Context(), buf)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) newTemplateData(r *http.Request) utils.TemplateData {
	return utils.TemplateData{
		CurrentYear: time.Now().Year(),
	}
}
