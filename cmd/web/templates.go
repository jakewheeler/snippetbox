package main

import (
	"snippetbox.jakewheeler.dev/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
