package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

var weathers [5]string = [5]string{"Sunny", "Rainy", "Smoggy", "Thundersnow", "Drizzly"}

func handler(w http.ResponseWriter, r *http.Request) {
	i := rand.Intn(4)
	temp := rand.Intn(120)
	w.Header().Set("Server", "WEATHER-MIDDLE")
	fmt.Fprintf(w, "%s and %d degrees F.", weathers[i], temp)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
