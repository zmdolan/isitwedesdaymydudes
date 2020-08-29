package main

import (
	"log"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		today := time.Now().Weekday()
		if today == time.Weekday(3) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "it is Wednesday my dudes"}`))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "it is not Wednesday my dudes"}`))
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "unsupported method"}`))
	}
}

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":5000", nil))
	//beanstalk forwards traffic to ec2 via port 5000, change to 8080 for localhost testing
}
