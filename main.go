package main

import (
	"net/http"

	"patrick.com/render-atl-hackathon/db"
)

func main() {
	database, err := db.GetDb()

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reviews, err := db.GetReviews(r.Response.Request.Context(), database)

		if err != nil {
			panic(err)
		}

	})
}
