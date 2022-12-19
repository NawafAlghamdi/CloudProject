package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

func HomeEndpoint(w http.ResponseWriter, req *http.Request) {
	lat1, _ := strconv.ParseFloat((req.URL.Query()["lat1"][0]), 64)
	lng1, _ := strconv.ParseFloat((req.URL.Query()["lng1"][0]), 64)
	lat2, _ := strconv.ParseFloat((req.URL.Query()["lat2"][0]), 64)
	lng2, _ := strconv.ParseFloat((req.URL.Query()["lng2"][0]), 64)
	fmt.Fprintf(w, "The distance is: %.2f Kilometers", distance(lat1, lng1, lat2, lng2, "K"))
}

func main() {
	http.HandleFunc("/", HomeEndpoint)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
