package views

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
)

// Template struct holds the parsed templates.
type Template struct {
	templates *template.Template
}

// NewTemplate initializes and returns a Template instance with all HTML templates parsed.
func NewTemplate() *Template {
	// Parse all HTML templates in the "templates" directory
	return &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}

// Render renders a template document.
// It implements echo.Renderer interface.
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Check if the template name is not empty
	if name == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "template name is empty")
	}

	// Execute the template with the provided data
	return t.templates.ExecuteTemplate(w, name, data)
}
