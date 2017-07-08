package grifts

import (
	"about_me/models"

	"github.com/markbates/grift/grift"
)

var _ = grift.Add("db:seed", func(c *grift.Context) error {
	user := &models.User{FullName: "Manuel Perez", Email: "mperez@wawand.co", Password: "manuel12345"}
	models.DB.Create(user)
	user = &models.User{FullName: "Edwin Polo", Email: "epolo@wawand.co", Password: "edwin12345"}
	models.DB.Create(user)
	user = &models.User{FullName: "Antonio Pagano", Email: "apagano@wawand.co", Password: "antonio12345"}
	models.DB.Create(user)
	user = &models.User{FullName: "Cristian Pelaez", Email: "cpelaez@wawand.co", Password: "cristian12345"}
	models.DB.Create(user)
	user = &models.User{FullName: "Jessica Villalobos", Email: "jvillalobos@wawand.co", Password: "12345"}
	models.DB.Create(user)

	return nil
})
