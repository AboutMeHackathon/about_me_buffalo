package actions

import (
	"about_me/models"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// UsersList default implementation.
func UsersList(c buffalo.Context) error {
	users := &models.Users{}
	tx := c.Get("tx").(*pop.Connection)
	err := tx.All(users)
	if err != nil {
		return c.Error(404, errors.WithStack(err))
	}

	c.Set("username", c.Session().Get("full_name"))
	c.Set("users", users)

	return c.Render(200, r.HTML("users/list.html"))
}

func UsersShowEntries(c buffalo.Context) error {
	userID := c.Param("user_id")
	tx := c.Value("tx").(*pop.Connection)
	entries := &models.Entries{}
	user := &models.User{}

	err := tx.Where("user_id = ?", userID).All(entries)
	if err != nil {
		return nil
	}

	err = tx.Find(user, userID)
	if err != nil {
		return nil
	}

	c.Set("username", user.FullName)
	c.Set("entries", entries)

	return c.Render(200, r.HTML("users/show_entries.html"))
}
