package actions

import (
	"about_me/models"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

func EntriesNew(c buffalo.Context) error {
	c.Set("entrie", models.Entrie{})
	c.Set("user_id", c.Session().Get("user_id"))
	return c.Render(200, r.HTML("entries/new.html"))
}

// EntriesCreate default implementation.
func EntriesCreate(c buffalo.Context) error {
	entrie := &models.Entrie{}

	err := c.Bind(entrie)
	if err != nil {
		return errors.WithStack(err)
	}

	tx := c.Get("tx").(*pop.Connection)
	err = tx.Create(entrie)

	if err != nil {
		return nil
	}

	return c.Redirect(301, "/")
}
