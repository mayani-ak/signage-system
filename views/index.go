package views

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// IndexView renders the index page for the Digital Signage system.
// It sends a title to be displayed on the page.
func IndexView(c echo.Context) error {
	// Data to be passed to the template
	data := map[string]interface{}{
		"title": "Digital Signage",
	}

	// Render the "index.html" template with the provided data
	return c.Render(http.StatusOK, "index.html", data)
}
