package main

import (
	"encoding/json"
	"expvar"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var weathers [5]string = [5]string{"Sunny", "Rainy", "Smoggy", "Thundersnow", "Drizzly"}

type Weather struct {
	Temp string
	Type string
}

func handler(w http.ResponseWriter, r *http.Request) {
	i := rand.Intn(4)
	temp := rand.Intn(120)
	weather := Weather{Temp: strconv.Itoa(temp) + "F", Type: weathers[i]}

	w.Header().Set("Server", "WEATHER")
	w.Header().Set("Content-Type", "application/json")

	jsonData, _ := json.Marshal(weather)
	w.Write(jsonData)
	fmt.Println("Dispatched weather data: ", weather)
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
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8070", nil)
}
