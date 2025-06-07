package utils

import (
	"time"

	"github.com/patvoj/snippetbox/internal/models"
)

type TemplateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []models.Snippet
}

func HumanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format("02 Jan 2006 at 15:04")
}
