# Hello!

## Links
https://jamboard.google.com/d/17opf6kCJIRen7uNHnfhRW5uAUh7mL8eOksG8zwGtuOA/viewer?ts=666b2db4&pli=1&f=0

https://github.com/newrelic-experimental/innovate-atl

## Running the app
You must set the `OPEN_API_TOKEN` environment variable to run the APP

SQLite create table statement:

```
CREATE TABLE feedback (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Review TEXT,
    OverallOpinion TEXT,

    MentionsDrivers BOOLEAN,
    OpinionOfDriver TEXT,
    DriversSummary TEXT,

    MentionsPurchasing BOOLEAN,
    OpinionOfPurchasing TEXT,
    PurchasingSummary TEXT,

    MentionsHomeless BOOLEAN,
    OpinionOfHomeless TEXT,
    HomelessSummary TEXT,

    MentionsAccessibility BOOLEAN,
    OpinionOfAccessibility TEXT,
    AccessibilitySummary TEXT,

    MentionsSafety BOOLEAN,
    OpinionOfSafety TEXT,
    SafetySummary TEXT,

    MentionsCustomerService BOOLEAN,
    OpinionOfCustomerService TEXT,
    CustomerServiceSummary TEXT,

    MentionsTime BOOLEAN,
    OpinionOfTime TEXT,
    TimeSummary TEXT,

    MentionsSignage BOOLEAN,
    OpinionOfSignage TEXT,
    SignageSummary TEXT,
    
    MentionsCleanliness BOOLEAN,
    OpinionOfCleanliness TEXT,
    CleanlinessSummary TEXT
);
```
