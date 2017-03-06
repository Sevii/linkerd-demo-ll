package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var weathers [5]string = [5]string{"Sunny", "Rainy", "Smoggy", "Thundersnow", "Drizzly"}

type status struct {
	Dow   int
	SP500 int
}

func handler(w http.ResponseWriter, r *http.Request) {
	dowlevel := (rand.Float64() * 24040)
	splevel := (rand.Float64() * 12834)
	s := status{Dow: int(dowlevel), SP500: int(splevel)}

	w.Header().Set("Server", "STOCKS")
	w.Header().Set("Content-Type", "application/json")

	jsonData, _ := json.Marshal(s)
	fmt.Println(string(jsonData))
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/", prometheus.InstrumentHandlerFunc("Stocks", handler))
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8075", nil)
}
