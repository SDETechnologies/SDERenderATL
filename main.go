package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"patrick.com/render-atl-hackathon/db"
)

func main() {
	/*
			reviewText := `Crazy crazy. My last Marta train.
		I’ve been taking Marta to airport from in town for a decade
		Since COVID the homeless people people have taken over. No policing.
		I’ve been threatened, exposed to fights of extreme verbal language and other inappropriate behavior. They - Marta- ignore all.
		A fellow passenger and I tonight were discussing how convenient it was but not worth it. She and several of us, of multiple races, were traumatized by the experience. No one cared.
		`
			review, err := chatgpt.SummarizeReview(context.Background(), reviewText)
			if err != nil {
				panic(err)
			}

			err = db.InsertReview(context.TODO(), database, review)

			if err != nil {
				panic(err)
			}

	*/
	database, err := db.GetDb()

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		overallStats, err := db.GetOverallRating(r.Context(), database)
		if err != nil {
			panic(err)
		}
		tmpl, err := template.ParseFiles("views/index.html")

		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(w, map[string]any{
			"OverallStats": overallStats,
		})
		if err != nil {
			panic(err)
		}

	})

	log.Fatal(http.ListenAndServe(":4040", nil))
}
