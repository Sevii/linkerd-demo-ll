package main

import (
	"encoding/json"
	"expvar"
	"fmt"
	"math/rand"
	"net/http"
	"time"
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
	go logExpvars()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8075", nil)
}

func logExpvars() {
	for true {
		expvar.Do(func(variable expvar.KeyValue) {
			fmt.Printf("expvar.Key: %s expvar.Value: %s \n", variable.Key, variable.Value)
		})
		time.Sleep(time.Second * 5)
	}
}
