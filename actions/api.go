package actions

import (
	"about_me/models"
	"encoding/json"
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

func ApiValidateLogin(c buffalo.Context) error {
	email := c.Request().Form.Get("Email")
	password := c.Request().Form.Get("Password")

	user, err := models.SearchUser(email, password)
	if err != nil {
		return nil
	}

	json.NewEncoder(c.Response()).Encode(user)
	return nil
}

func ApiGetEntries(c buffalo.Context) error {
	userID := c.Param("user_id")

	tx := c.Value("tx").(*pop.Connection)
	entries := &models.Entries{}

	err := tx.Where("user_id = ?", userID).All(entries)
	if err != nil {
		return nil
	}

	fmt.Println("entries")
	fmt.Println(entries)

	json.NewEncoder(c.Response()).Encode(entries)
	return nil
}

func ApiUsersList(c buffalo.Context) error {
	users := &models.Users{}
	tx := c.Get("tx").(*pop.Connection)
	err := tx.All(users)
	if err != nil {
		return c.Error(404, errors.WithStack(err))
	}

	json.NewEncoder(c.Response()).Encode(users)
	return nil
}

func ApiGetEntriesByUserID(c buffalo.Context) error {
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

	result := map[string]interface{}{
		"entries": entries,
		"friend":  user.FullName,
	}
	json.NewEncoder(c.Response()).Encode(result)
	return nil
}
