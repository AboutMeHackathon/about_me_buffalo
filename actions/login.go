package actions

import (
	"about_me/models"

	"github.com/gobuffalo/buffalo"
)

func LoginMW(h buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if "/session/new" == c.Request().URL.Path {
			return h(c)
		}

		if c.Session().Get("user_id") == nil {
			return c.Redirect(302, "/session/new")
		}

		c.Set("user", map[string]interface{}{
			"full_name": c.Session().Get("full_name"),
			"user_id":   c.Session().Get("user_id"),
		})

		return h(c)
	}
}

func ValidateLogin(c buffalo.Context) error {
	email := c.Request().Form.Get("Email")
	password := c.Request().Form.Get("Password")

	user, err := models.SearchUser(email, password)
	if err != nil {
		c.Flash().Add("error", "Invalid Email or Password")
		return c.Redirect(301, "/session/new")
	}

	c.Session().Set("user_id", user.ID)
	c.Session().Set("full_name", user.FullName)
	c.Session().Save()

	return c.Redirect(301, "/")
}

// LoginNew default implementation.
func LoginNew(c buffalo.Context) error {
	c.Set("user", models.User{})
	return c.Render(200, r.HTML("login/new.html"))
}

func Logout(c buffalo.Context) error {
	c.Session().Clear()
	c.Session().Save()

	return c.Redirect(301, "/session/new")
}
