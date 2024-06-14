package controllers

import (
	"github.com/labstack/echo/v4"
	"patrick.com/render-atl-hackathon/db"
)

func IndexPage(c echo.Context) error {
	database, err := db.GetDb()

	if err != nil {
		panic(err)
	}

	reviews, err := db.GetReviews(c.Request().Context(), database)

	if err != nil {
		panic(err)
	}

	return c.Render(200, "index.html", map[string]any{
		"Reviews": reviews,
	})
}
