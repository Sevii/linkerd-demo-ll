package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"expvar"
)

const (
	weatherUrl = "http://weather"
	stocksUrl  = "http://stocks"
)

type weather struct {
	Temp string
	Type string
}

type stocks struct {
	Dow   int
	SP500 int
}

type user struct {
	Level    string
	Username string
}

type Report struct {
	Username string
	Level    string
	Dow      int
	SP500    int
	Temp     string
	Weather  string
	Time     string
}

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

// Get Weather
func getStocks() (stocks, error) {
	stocksResponse, err := netClient.Get(stocksUrl)
	if err != nil {
		return stocks{}, err
	}
	var s stocks
	err = json.NewDecoder(stocksResponse.Body).Decode(&s)
	if err != nil {
		return stocks{}, err
	}
	fmt.Println("Got stocks report: ", s)
	return s, nil
}

// Get Weather
func getWeather() (weather, error) {
	weatherResponse, err := netClient.Get(weatherUrl)
	if err != nil {
		return weather{}, err
	}
	var w weather
	err = json.NewDecoder(weatherResponse.Body).Decode(&w)
	if err != nil {
		return weather{}, err
	}
	fmt.Println("Got weather report: ", w)
	return w, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "REPORTS")
	var u user
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	weatherData, err := getWeather()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	s, err := getStocks()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	today := fmt.Sprintf(time.Now().Format(time.RFC850))

	report := Report{Username: u.Username, Level: u.Level, Dow: s.Dow, SP500: s.SP500, Temp: weatherData.Temp, Weather: weatherData.Type, Time: today}

	fmt.Println("Created report: ", report)

	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(report)
	w.Write(jsonData)
}

func logExpvars() {
	for true {
		expvar.Do(func(variable expvar.KeyValue) {
			fmt.Printf("expvar.Key: %s expvar.Value: %s \n", variable.Key, variable.Value)
		})
		time.Sleep(time.Second * 5)
	}
}

func main() {
	go logExpvars()
	http.HandleFunc("/report", handler)
	http.ListenAndServe(":8055", nil)
}