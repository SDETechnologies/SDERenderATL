package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"patrick.com/render-atl-hackathon/db"
)

func main() {

	database, err := db.GetDb()

	if err != nil {
		panic(err)
	}

	/*
	f, err := os.ReadFile("./reviews.json")
	if err != nil { panic(err) }

	err = json.Unmarshal(f, &reviews)
	if err != nil {panic(err)}

	fmt.Println("total reviews", len(reviews))
	for i, r := range reviews[10:] {

		err = json.Unmarshal(f, &reviews)
		if err != nil { panic(err)}
		review, err := chatgpt.SummarizeReview(context.Background(), r)
		if err != nil {
			panic(err)
		}

		err = db.InsertReview(context.TODO(), database, review)

		if err != nil {
			panic(err)
		}
	}
	*/

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("views/index.html")

		if err != nil {
			panic(err)
		}

		overallStats, err := db.GetOverallRating(r.Context(), database)
		if err != nil {
			panic(err)
		}

		mentionTopicCounts, err := db.GetNumberOfTopicsCount(r.Context(), database)
		if err != nil {
			panic(err)
		}

		sum, err := db.GetSumOpinionsOfTopic(r.Context(), database)
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(w, map[string]any{
			"OverallStats": overallStats,
			"TopicCounts":  mentionTopicCounts,
			"SumOpinions":  sum,
		})

		if err != nil {
			panic(err)
		}

	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	port := "80"
	if os.Getenv("HACKATHON_PORT") != "" {
		port = os.Getenv("HACKATHON_PORT")
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
