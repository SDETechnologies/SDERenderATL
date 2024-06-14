
package reviews

import (
	"fmt"
    "os"
	// "io"
	"github.com/go-rod/rod"
	// "github.com/go-rod/rod/lib/launcher"
	"github.com/joho/godotenv"
	// "github.com/JustinBeckwith/go-yelp"
	// "encoding/json"
	"googlemaps.github.io/maps"
    "context"
)

type Station struct {
    Name string `json:"name"`
    Date string `json:"date"`
    Review string `json:"review"`
}

type ScrapeStationResponse struct {
    Name string `json:"name"`
    Date string `json:"date"`
    Reviews string `json:"review"`
}

type Review struct {
    Text string
}


func getReviewFromStationName(stationName string) string {
    nameForAPI := stationName + " Indian Creek MARTA Station"
    return nameForAPI
}


func getStationReview(ctx context.Context, stationName string) ([]string,error) {

    endReviews := []string{} 

    c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))
    if err != nil {
        return endReviews, fmt.Errorf("Error querying google api: %s", err)
    }

    request := maps.FindPlaceFromTextRequest {
        Input: stationName,
        InputType: "textquery",
    }
    // fmt.Println("request: ", request)

    response,err := c.FindPlaceFromText(ctx, &request)
    if err != nil {
        return endReviews, fmt.Errorf("Error querying google api: %s", err)
    }
    // fmt.Println("response: ", response.PlaceID)
    // fmt.Println("response: ", response.Candidates)

    fmt.Println("\n----------------------------------------------------------------\n")


    if len(response.Candidates ) == 0 {
        return  []string{}, nil
    }
    placeID := response.Candidates[0].PlaceID

    placeDetailsRequest := maps.PlaceDetailsRequest{
        PlaceID: string(placeID),
        Language: "en",
    }
    res, err := c.PlaceDetails(ctx, &placeDetailsRequest)
    if err != nil {
        panic(err)
    }
    for _, review := range res.Reviews {

        endReviews = append(endReviews, review.Text)
    }

    return endReviews,  nil
}


// // func scrapeStationNames (w http.ResponseWriter, r *http.Request) {
// func scrapeStationNames () ([]Station, error) {
//     stations := []Station{}
//     browser := rod.New().MustConnect()
//     page := *browser.MustPage("https://www.itsmarta.com/train-stations-and-schedules.aspx").MustWaitLoad()
//     stationElementsDiv := page.MustElement("stations__items isotope")
//     fmt.Println("stationElementsDiv: ", stationElementsDiv)
//     stationElements := stationElementsDiv.MustElements("stations__item")
//     fmt.Println("stationElements: ", stationElements)
//
//     for _,stationElement := range stationElements {
//         fmt.Println("stationElement: ", stationElement)
//         station := Station{}
//         fmt.Println("station: ", station)
//         stationName := stationElement.MustElem
//         ent("stations__item-name").MustText()
//         station.Name = stationName
//         stations = append(stations, station)
//     } 
//     return stations, nil
// }

func scrapeStationNames2 (ctx context.Context) ([]string, error) {
    stations := []string{}
    browser := *rod.New().MustConnect()
    page := browser.MustPage("https://www.itsmarta.com/train-stations-and-schedules.aspx").MustWaitLoad()
    stationElements := page.MustElements("a[class='stations__item-name'")

    for _,stationElement := range stationElements {
        stationName := stationElement.MustText()
        // stationReview := getStationReview(ctx, stationName) 

        stationReview,err := getStationReview(context.Background(), stationName)
        if err != nil { panic(err) }
        stations = append(stations, stationReview...)
        fmt.Println("stationReview: ", stationReview)

    } 
    return stations, nil
}

const staticDir string = "/static/"

func GetAllReviews() []string{
    godotenv.Load()
    // port := os.Getenv("PORT")

    // router := router.NewRouter(service.NewService(database.NewDatabase(database.GetDB())))

    // r.HandleFunc("/getstationrowelement", scrapeStationNames)
    //
    // res,err := http.Get("/getstationrowelement")
    // if err != nil {panic(err)}
    //
    // data,err := io.ReadAll(res.Body)
    // if err != nil {panic(err)}
    //
    //
    //
    // testResponse := ScrapeStationResponse{}
    // err = json.Unmarshal(data, &testResponse)
    // fmt.Println("testResponse: ", testResponse)

    stations, err  := scrapeStationNames2(context.TODO())
    if err != nil { panic(err)}
return stations

}

