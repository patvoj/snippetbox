package types

import "github.com/patvoj/snippetbox/internal/models"

type TemplateData struct {
	Snippet  *models.Snippet
	Snippets []models.Snippet
}
