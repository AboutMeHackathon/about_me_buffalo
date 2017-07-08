package actions

import (
	"about_me/models"
	"encoding/json"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// EntriesNew default implementation.
func EntriesNew(c buffalo.Context) error {
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

	json.NewEncoder(c.Response()).Encode(entrie)

	// return nil
	return c.Render(200, r.HTML("entries/new.html"))
}
