package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const JetstreamApiUrl = "https://api.jetstreamrfid.com/3/events/50" // Jetstream API URL that calls for 50 events, *will mock this response

type Parameter struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type SensorReading struct {
	Name        string `json:"Name"`
	Value       string `json:"Value"`
	ReadingTime string `json:"ReadingTime"`
}

type Event struct {
	CommandID           string          `json:"CommandId,omitempty"`
	CommandName         string          `json:"CommandName,omitempty"`
	OutputParameterList []string        `json:"OutputParameterList,omitempty"`
	ExceptionList       []string        `json:"ExceptionList,omitempty"`
	Device              string          `json:"Device"`
	ReceivedTime        string          `json:"ReceivedTime,omitempty"`
	Type                string          `json:"Type"`
	EventID             string          `json:"EventId"`
	EventTime           string          `json:"EventTime"`
	Version             int             `json:"Version"`
	URI                 string          `json:"URI,omitempty"`
	Verb                string          `json:"Verb,omitempty"`
	User                string          `json:"User,omitempty"`
	Parameters          []Parameter     `json:"Parameters,omitempty"`
	SensorReadings      []SensorReading `json:"SensorReadings,omitempty"`
}

type EventsResponse struct {
	BatchId string  `json:"Batchid"`
	Count   int     `json:"Count"`
	Events  []Event `json:"Events"`
}

type MockJetstream struct{}

func handleGetEvents(w http.ResponseWriter, r *http.Request) {
	jetstreamJsonRes := EventsResponse{ // assuming we make http request to JetstreamApiUrl and recieve events response
		BatchId: "c15d316b-5fc3-4a7f-a2ba-b1a96aabcd32",
		Count:   5,
		Events: []Event{
			{
				CommandID:           "2bf00848-b06a-4820-86fc-58c12e08ee3e",
				CommandName:         "SetConfigValuesCommand",
				OutputParameterList: []string{},
				ExceptionList:       []string{},
				Device:              "MyDeviceName1",
				ReceivedTime:        "2018-05-19T01:16:23Z",
				Type:                "CommandCompletionEvent",
				EventID:             "93723c89-cb44-44f0-8d86-87ede0fb3ba6",
				EventTime:           "2018-05-19T01:16:19Z",
				Version:             2,
			},
			{
				Device:      "MyDeviceName1",
				CommandID:   "2bf00848-b06a-4820-86fc-58c12e08ee3e",
				CommandName: "SetConfigValuesCommand",
				URI:         "/2/Devices/MyDeviceName1/policy/sync",
				Verb:        "Post",
				User:        "MyUserName",
				Parameters: []Parameter{
					{
						Key:   "aggregateeventscancount",
						Value: "4",
					},
				},
				Type:      "CommandQueuedEvent",
				EventID:   "f3df515b-c53a-4b3b-b3e9-01ffe6ffd8f3",
				EventTime: "2018-05-19T01:16:21Z",
				Version:   3,
			},
			{
				Device:       "MyDeviceName2",
				ReceivedTime: "2018-05-19T01:19:44Z",
				Type:         "HeartbeatEvent",
				EventID:      "b55cf245-4a17-438e-8b12-538c3ef6f536",
				EventTime:    "2018-05-19T01:19:43Z",
				Version:      3,
			},
			{
				Device:       "MyDeviceName1",
				ReceivedTime: "2018-05-19T02:05:40Z",
				Type:         "HeartbeatEvent",
				EventID:      "9e5d2294-655e-4a11-ba02-e504b435e75f",
				EventTime:    "2018-05-19T02:05:36Z",
				Version:      3,
			},
			{
				SensorReadings: []SensorReading{
					{
						Name:        "TemperatureA",
						Value:       "-16",
						ReadingTime: "2018-05-19T02:06:25Z",
					},
					{
						Name:        "TemperatureB",
						Value:       "-61",
						ReadingTime: "2018-05-19T02:06:25Z",
					},
				},
				Device:       "MyDeviceName1",
				ReceivedTime: "2018-05-19T02:06:28Z",
				Type:         "SensorReadingEvent",
				EventID:      "66524199-0c71-40d2-a36a-c05983b70355",
				EventTime:    "2018-05-19T02:06:25Z",
				Version:      3,
			},
		},
	}

	eventTypeCount := map[string]int{} // map to store count of each event type

	for _, event := range jetstreamJsonRes.Events { // looping through events response and counting each event type
		eventTypeCount[event.Type]++
	}

	eventsCountJson, err := json.Marshal(eventTypeCount)

	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(eventsCountJson) // writing number of each unique event type to client
	}

	// fmt.Print(jetstreamJsonRes) // printing events response
}

func requestHandler() { // handler function to wrap API routes
	http.HandleFunc("/", handleGetEvents) // mapping root '/' endpoint to eventsPage func
	http.ListenAndServe(":3001", nil)
}

func main() {
	requestHandler() // calling api request handler function
}
