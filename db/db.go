package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Review struct {
	Id      int
	Content string
}

func GetDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return nil, fmt.Errorf("[DB] opening db: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("[DB] pinging db: %s", err)
	}

	return db, nil
}

func InsertReview(ctx context.Context, db *sql.DB, feedback Feedback) error {
	sqlStmt := "INSERT INTO feedback(Review, OverallOpinion, MentionsDrivers, OpinionOfDriver, DriversSummary, MentionsPurchasing, OpinionOfPurchasing, PurchasingSummary, MentionsHomeless, OpinionOfHomeless, HomelessSummary, MentionsAccessibility, OpinionOfAccessibility, AccessibilitySummary, MentionsSafety, OpinionOfSafety, SafetySummary, MentionsCustomerService, OpinionOfCustomerService, CustomerServiceSummary, MentionsTime, OpinionOfTime, TimeSummary, MentionsSignage, OpinionOfSignage, SignageSummary, MentionsCleanliness, OpinionOfCleanliness, CleanlinessSummary) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("[DB] beginning transaction: %s", err)
	}

	stmt, err := tx.Prepare(sqlStmt)

	if err != nil {
		return fmt.Errorf("[DB] preparing statement: %s", err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, feedback.Review, feedback.OverallOpinion, feedback.MentionsDrivers, feedback.OpinionOfDriver, feedback.DriversSummary, feedback.MentionsPurchasing, feedback.OpinionOfPurchasing, feedback.PurchasingSummary, feedback.MentionsHomeless, feedback.OpinionOfHomeless, feedback.HomelessSummary, feedback.MentionsAccessibility, feedback.OpinionOfAccessibility, feedback.AccessibilitySummary, feedback.MentionsSafety, feedback.OpinionOfSafety, feedback.SafetySummary, feedback.MentionsCustomerService, feedback.OpinionOfCustomerService, feedback.CustomerServiceSummary, feedback.MentionsTime, feedback.OpinionOfTime, feedback.TimeSummary, feedback.MentionsSignage, feedback.OpinionOfSignage, feedback.SignageSummary, feedback.MentionsCleanliness, feedback.OpinionOfCleanliness, feedback.CleanlinessSummary)

	if err != nil {
		return fmt.Errorf("[DB] executing statement: %s", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("[DB] committing transcation: %s", err)
	}

	return nil
}

type Feedback struct {
	Review         string
	OverallOpinion string `json:"overallOpinion"`

	MentionsDrivers bool   `json:"mentionsDrivers"`
	OpinionOfDriver string `json:"opinionOfDriver"`
	DriversSummary  string `json:"driversSummary"`

	MentionsPurchasing  bool   `json:"mentionsPurchasing"`
	OpinionOfPurchasing string `json:"opinionOfPurchasing"`
	PurchasingSummary   string `json:"purchasingSummary"`

	MentionsHomeless  bool   `json:"mentionsHomeless"`
	OpinionOfHomeless string `json:"opinionOfHomeless"`
	HomelessSummary   string `json:"homelessSummary"`

	MentionsAccessibility  bool   `json:"mentionsAccessibility"`
	OpinionOfAccessibility string `json:"opinionOfAccessibility"`
	AccessibilitySummary   string `json:"accessibilitySummary"`

	MentionsSafety  bool   `json:"mentionsSafety"`
	OpinionOfSafety string `json:"opinionOfSafety"`
	SafetySummary   string `json:"safetySummary"`

	MentionsCustomerService  bool   `json:"mentionsCustomerService"`
	OpinionOfCustomerService string `json:"opinionOfCustomerService"`
	CustomerServiceSummary   string `json:"customerServiceSummary"`

	MentionsTime  bool   `json:"mentionsTime"`
	OpinionOfTime string `json:"opinionOfTime"`
	TimeSummary   string `json:"timeSummary"`

	MentionsSignage  bool   `json:"mentionsSignage"`
	OpinionOfSignage string `json:"opinionOfSignage"`
	SignageSummary   string `json:"signageSummary"`

	MentionsCleanliness  bool   `json:"mentionsCleanliness"`
	OpinionOfCleanliness string `json:"opinionOfCleanliness"`
	CleanlinessSummary   string `json:"cleanlinessSummary"`
}

type OverallRating struct {
	Average          int
	Positive         int
	SlightlyPositive int
	Mixed            int
	SlightlyNegative int
	Negative         int
	Stars            []string
}

func GetOverallRating(ctx context.Context, db *sql.DB) (OverallRating, error) {
	sqlStmt := "select OverallOpinion, count(*) from feedback group by OverallOpinion"

	tx, err := db.Begin()
	if err != nil {
		return OverallRating{}, fmt.Errorf("[DB] beginning transaction: %s", err)
	}

	stmt, err := tx.Prepare(sqlStmt)

	if err != nil {
		return OverallRating{}, fmt.Errorf("[DB] preparing statement: %s", err)
	}

	defer stmt.Close()

	res, err := stmt.QueryContext(ctx)

	if err != nil {
		return OverallRating{}, fmt.Errorf("[DB] executing statement: %s", err)
	}

	overallRating := OverallRating{
		Stars: []string{},
	}

	for res.Next() {
		var opinion string
		var count int
		err = res.Scan(&opinion, &count)

		if err != nil {
			return OverallRating{}, fmt.Errorf("[DB] scanning thread")
		}

		switch opinion {
		case "Positive":
			overallRating.Positive = count
			overallRating.Average += 5
		case "Slightly Positive":
			overallRating.SlightlyPositive = count
			overallRating.Average += 4
		case "Mixed":
			overallRating.Mixed = count
			overallRating.Average += 3
		case "Slightly Negative":
			overallRating.SlightlyNegative = count
			overallRating.Average += 2
		case "Negative":
			overallRating.Negative = count
			overallRating.Average += 1
		}

	}
	overallRating.Average = overallRating.Average / 5

	for i := 0; i < 5; i++ {
		if i < overallRating.Average {
			overallRating.Stars = append(overallRating.Stars, "Full")
		} else if i == overallRating.Average {
			overallRating.Stars = append(overallRating.Stars, "HalfFull")
		} else {
			overallRating.Stars = append(overallRating.Stars, "Empty")
		}
	}

	return overallRating, nil
}
