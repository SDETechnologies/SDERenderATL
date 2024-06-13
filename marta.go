package main

/*
import (


type Train struct {
	Destination    string `json:"DESTINATION"`
	Direction      string `json:"DIRECTION"`
	EventTime      string `json:"EVENT_TIME"`
	Line           string `json:"LINE"`
	NextArr        string `json:"NEXT_ARR"`
	Station        string `json:"STATION"`
	TrainID        string `json:"TRAIN_ID"`
	WaitingSeconds string `json:"WAITING_SECONDS"`
	WaitingTime    string `json:"WAITING_TIME"`
}

const (
	BaseMartaURL  = "https://developer.itsmarta.com"
	RealtimeTrain = "/RealtimeTrain/RestServiceNextTrain/GetRealtimeArrivals"
)

func GetTrains() ([]Train, error) {
	res, err := http.Get(fmt.Sprintf("%s%s", BaseMartaURL, RealtimeTrain))

	if err != nil {
		return nil, fmt.Errorf("performing getting trains: %s", err)
	}

	realTimeTrains := []Train{}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("reading getting trains body: %s", err)
	}

	err = json.Unmarshal(body, &realTimeTrains)

	if err != nil {
		return nil, fmt.Errorf("unmarshaling trains data: %s", err)
	}

	return realTimeTrains, nil
}
*/
