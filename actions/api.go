package actions

import (
	"about_me/models"
	"encoding/json"
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
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
