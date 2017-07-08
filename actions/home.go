package actions

import (
	"about_me/models"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	userID := c.Session().Get("user_id")
	tx := c.Value("tx").(*pop.Connection)
	entries := &models.Entries{}

	err := tx.Where("user_id = ?", userID).All(entries)
	if err != nil {
		return nil
	}

	c.Set("entries", entries)
	c.Set("username", c.Session().Get("full_name"))

	return c.Render(200, r.HTML("index.html"))
}
