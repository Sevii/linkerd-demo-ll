package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
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

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8070", nil)
}
