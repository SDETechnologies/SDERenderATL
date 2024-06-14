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

type MentionsTopicsCounts struct {
	MentionsDriversCount         int
	MentionsPurchasingCount      int
	MentionsHomelessCount        int
	MentionsAccessibilityCount   int
	MentionsSafetyCount          int
	MentionsCustomerServiceCount int
	MentionsTimeCount            int
	MentionsSignageCount         int
	MentionsCleanlinessCount     int
}

func GetNumberOfTopicsCount(ctx context.Context, db *sql.DB) (MentionsTopicsCounts, error) {
	sqlStmt := `
	SELECT
    SUM(CASE WHEN MentionsDrivers = 1 THEN 1 ELSE 0 END) AS MentionsDriversCount,
    SUM(CASE WHEN MentionsPurchasing = 1 THEN 1 ELSE 0 END) AS MentionsPurchasingCount,
    SUM(CASE WHEN MentionsHomeless = 1 THEN 1 ELSE 0 END) AS MentionsHomelessCount,
    SUM(CASE WHEN MentionsAccessibility = 1 THEN 1 ELSE 0 END) AS MentionsAccessibilityCount,
    SUM(CASE WHEN MentionsSafety = 1 THEN 1 ELSE 0 END) AS MentionsSafetyCount,
    SUM(CASE WHEN MentionsCustomerService = 1 THEN 1 ELSE 0 END) AS MentionsCustomerServiceCount,
    SUM(CASE WHEN MentionsTime = 1 THEN 1 ELSE 0 END) AS MentionsTimeCount,
    SUM(CASE WHEN MentionsSignage = 1 THEN 1 ELSE 0 END) AS MentionsSignageCount,
    SUM(CASE WHEN MentionsCleanliness = 1 THEN 1 ELSE 0 END) AS MentionsCleanlinessCount
FROM feedback;
	`
	tx, err := db.Begin()
	if err != nil {
		return MentionsTopicsCounts{}, fmt.Errorf("[DB] beginning transaction: %s", err)
	}

	stmt, err := tx.Prepare(sqlStmt)

	if err != nil {
		return MentionsTopicsCounts{}, fmt.Errorf("[DB] preparing statement: %s", err)
	}

	defer stmt.Close()

	res, err := stmt.QueryContext(ctx)

	if err != nil {
		return MentionsTopicsCounts{}, fmt.Errorf("[DB] executing statement: %s", err)
	}

	overallRating := MentionsTopicsCounts{}

	for res.Next() {
		err = res.Scan(&overallRating.MentionsDriversCount, &overallRating.MentionsPurchasingCount, &overallRating.MentionsHomelessCount, &overallRating.MentionsAccessibilityCount, &overallRating.MentionsSafetyCount, &overallRating.MentionsCustomerServiceCount, &overallRating.MentionsTimeCount, &overallRating.MentionsSignageCount, &overallRating.MentionsCleanlinessCount)

		if err != nil {
			return MentionsTopicsCounts{}, fmt.Errorf("[DB] scanning: %s", err)
		}

	}
	return overallRating, nil
}

type OpinionsSummary struct {
	SumDriverPositive          int
	SumDriverMixed             int
	SumDriverNegative          int
	SumPurchasingPositive      int
	SumPurchasingMixed         int
	SumPurchasingNegative      int
	SumHomelessPositive        int
	SumHomelessMixed           int
	SumHomelessNegative        int
	SumAccessibilityPositive   int
	SumAccessibilityMixed      int
	SumAccessibilityNegative   int
	SumSafetyPositive          int
	SumSafetyMixed             int
	SumSafetyNegative          int
	SumCustomerServicePositive int
	SumCustomerServiceMixed    int
	SumCustomerServiceNegative int
	SumTimePositive            int
	SumTimeMixed               int
	SumTimeNegative            int
	SumSignagePositive         int
	SumSignageMixed            int
	SumSignageNegative         int
	SumCleanlinessPositive     int
	SumCleanlinessMixed        int
	SumCleanlinessNegative     int
}

func GetSumOpinionsOfTopic(ctx context.Context, db *sql.DB) (OpinionsSummary, error) {
	query := `SELECT
    SUM(CASE WHEN OpinionOfDriver = 'Positive' THEN 1 ELSE 0 END) AS SumDriverPositive,
    SUM(CASE WHEN OpinionOfDriver = 'Mixed' THEN 1 ELSE 0 END) AS SumDriverMixed,
    SUM(CASE WHEN OpinionOfDriver = 'Negative' THEN 1 ELSE 0 END) AS SumDriverNegative,
    
    SUM(CASE WHEN OpinionOfPurchasing = 'Positive' THEN 1 ELSE 0 END) AS SumPurchasingPositive,
    SUM(CASE WHEN OpinionOfPurchasing = 'Mixed' THEN 1 ELSE 0 END) AS SumPurchasingMixed,
    SUM(CASE WHEN OpinionOfPurchasing = 'Negative' THEN 1 ELSE 0 END) AS SumPurchasingNegative,

    SUM(CASE WHEN OpinionOfHomeless = 'Positive' THEN 1 ELSE 0 END) AS SumHomelessPositive,
    SUM(CASE WHEN OpinionOfHomeless = 'Mixed' THEN 1 ELSE 0 END) AS SumHomelessMixed,
    SUM(CASE WHEN OpinionOfHomeless = 'Negative' THEN 1 ELSE 0 END) AS SumHomelessNegative,

    SUM(CASE WHEN OpinionOfAccessibility = 'Positive' THEN 1 ELSE 0 END) AS SumAccessibilityPositive,
    SUM(CASE WHEN OpinionOfAccessibility = 'Mixed' THEN 1 ELSE 0 END) AS SumAccessibilityMixed,
    SUM(CASE WHEN OpinionOfAccessibility = 'Negative' THEN 1 ELSE 0 END) AS SumAccessibilityNegative,

    SUM(CASE WHEN OpinionOfSafety = 'Positive' THEN 1 ELSE 0 END) AS SumSafetyPositive,
    SUM(CASE WHEN OpinionOfSafety = 'Mixed' THEN 1 ELSE 0 END) AS SumSafetyMixed,
    SUM(CASE WHEN OpinionOfSafety = 'Negative' THEN 1 ELSE 0 END) AS SumSafetyNegative,

    SUM(CASE WHEN OpinionOfCustomerService = 'Positive' THEN 1 ELSE 0 END) AS SumCustomerServicePositive,
    SUM(CASE WHEN OpinionOfCustomerService = 'Mixed' THEN 1 ELSE 0 END) AS SumCustomerServiceMixed,
    SUM(CASE WHEN OpinionOfCustomerService = 'Negative' THEN 1 ELSE 0 END) AS SumCustomerServiceNegative,

    SUM(CASE WHEN OpinionOfTime = 'Positive' THEN 1 ELSE 0 END) AS SumTimePositive,
    SUM(CASE WHEN OpinionOfTime = 'Mixed' THEN 1 ELSE 0 END) AS SumTimeMixed,
    SUM(CASE WHEN OpinionOfTime = 'Negative' THEN 1 ELSE 0 END) AS SumTimeNegative,

    SUM(CASE WHEN OpinionOfSignage = 'Positive' THEN 1 ELSE 0 END) AS SumSignagePositive,
    SUM(CASE WHEN OpinionOfSignage = 'Mixed' THEN 1 ELSE 0 END) AS SumSignageMixed,
    SUM(CASE WHEN OpinionOfSignage = 'Negative' THEN 1 ELSE 0 END) AS SumSignageNegative,

    SUM(CASE WHEN OpinionOfCleanliness = 'Positive' THEN 1 ELSE 0 END) AS SumCleanlinessPositive,
    SUM(CASE WHEN OpinionOfCleanliness = 'Mixed' THEN 1 ELSE 0 END) AS SumCleanlinessMixed,
    SUM(CASE WHEN OpinionOfCleanliness = 'Negative' THEN 1 ELSE 0 END) AS SumCleanlinessNegative
FROM
    feedback;
`
	tx, err := db.Begin()
	if err != nil {
		return OpinionsSummary{}, fmt.Errorf("[DB] beginning transaction: %s", err)
	}

	stmt, err := tx.Prepare(query)

	if err != nil {
		return OpinionsSummary{}, fmt.Errorf("[DB] preparing statement: %s", err)
	}

	defer stmt.Close()

	res, err := stmt.QueryContext(ctx)

	if err != nil {
		return OpinionsSummary{}, fmt.Errorf("[DB] executing statement: %s", err)
	}

	opinionsSummary := OpinionsSummary{}

	for res.Next() {
		err = res.Scan(
			&opinionsSummary.SumDriverPositive,
			&opinionsSummary.SumDriverMixed,
			&opinionsSummary.SumDriverNegative,
			&opinionsSummary.SumPurchasingPositive,
			&opinionsSummary.SumPurchasingMixed,
			&opinionsSummary.SumPurchasingNegative,
			&opinionsSummary.SumHomelessPositive,
			&opinionsSummary.SumHomelessMixed,
			&opinionsSummary.SumHomelessNegative,
			&opinionsSummary.SumAccessibilityPositive,
			&opinionsSummary.SumAccessibilityMixed,
			&opinionsSummary.SumAccessibilityNegative,
			&opinionsSummary.SumSafetyPositive,
			&opinionsSummary.SumSafetyMixed,
			&opinionsSummary.SumSafetyNegative,
			&opinionsSummary.SumCustomerServicePositive,
			&opinionsSummary.SumCustomerServiceMixed,
			&opinionsSummary.SumCustomerServiceNegative,
			&opinionsSummary.SumTimePositive,
			&opinionsSummary.SumTimeMixed,
			&opinionsSummary.SumTimeNegative,
			&opinionsSummary.SumSignagePositive,
			&opinionsSummary.SumSignageMixed,
			&opinionsSummary.SumSignageNegative,
			&opinionsSummary.SumCleanlinessPositive,
			&opinionsSummary.SumCleanlinessMixed,
			&opinionsSummary.SumCleanlinessNegative,
		)
		if err != nil {
			return OpinionsSummary{}, fmt.Errorf("[DB] scanning: %s", err)
		}

	}
	return opinionsSummary, nil

}
