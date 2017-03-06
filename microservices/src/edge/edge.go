package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"expvar"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

const (
	loginUrl   = "http://login/login"
	reportsUrl = "http://reports/report"
)

type user struct {
	Level    string
	Username string
}

type userLogin struct {
	Password string
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

var templates = template.Must(template.ParseFiles("report.html"))

// Get report
func getReport(u user) (Report, error) {
	jsonData, _ := json.Marshal(u)
	jsonArr := []byte(jsonData)
	resp, err := netClient.Post(reportsUrl, "application/json", bytes.NewBuffer(jsonArr))
	if err != nil {
		return Report{}, err
	}
	defer resp.Body.Close()

	var r Report
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return Report{}, err
	}
	fmt.Println("Got  report: ", r)
	return r, nil
}

// Check user login
func getLogin(u userLogin) (user, error) {
	jsonData, _ := json.Marshal(u)
	jsonArr := []byte(jsonData)
	fmt.Println("User login: ", string(jsonData))
	resp, err := netClient.Post(loginUrl, "application/json", bytes.NewBuffer(jsonArr))

	if err != nil {
		fmt.Println("Got an error: ", err)
		return user{}, err
	}
	var userData user
	err = json.NewDecoder(resp.Body).Decode(&userData)
	if err != nil {
		return user{}, err
	}
	fmt.Println("Got user data: ", userData)
	return userData, nil
}
func checkAuth(w http.ResponseWriter, r *http.Request) (user, error) {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return user{}, errors.New("No login data found in headers")
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return user{}, errors.New("Failed to decode auth")
	}
	fmt.Println("decoded: ", b)

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return user{}, errors.New("Auth unreadable")
	}
	userName := pair[0]
	password := pair[1]
	fmt.Println("Password: %s, Username: %s \n", userName, password)
	return getLogin(userLogin{Password: password, Username: userName})
}

func handler(w http.ResponseWriter, r *http.Request) {
	u, err := checkAuth(w, r)
	if err != nil {
		w.Header().Set("WWW-Authenticate", `Basic realm="demo.realm`)
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
		fmt.Println("Login failed, ", r)
		return
	}
	fmt.Println("Logged in; ", u)
	//get report
	report, err := getReport(u)
	if err != nil {
		fmt.Println("Failed to get report ", err)
		return
	}
	fmt.Println(report)
	err = templates.ExecuteTemplate(w, "report.html", report)

}

func main() {
	go logExpvars()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func logExpvars() {
	for true {
		expvar.Do(func(variable expvar.KeyValue) {
			fmt.Println("expvar.Key: %s expvar.Value: %s \n", variable.Key, variable.Value)
		})
		time.Sleep(time.Second * 5)
	}
}
