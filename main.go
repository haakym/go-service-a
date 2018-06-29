package main

import (
	"encoding/json"
	"net/http"
)

type Coupon struct {
	Brand string `json:"brand"`
	Value int    `json:"value"`
}

type Coupons []Coupon

func main() {
	http.HandleFunc("/coupons", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(response http.ResponseWriter, request *http.Request) {

	/*
		Validate request
		if invalid, respond with helpful message
		parameters := request.URL.Query()
		log.Println(parameters)
	*/

	/*
		Authentication? JWT?
		Request to Service B
			Parse response
			resp, err := http.Get("http://localhost:3000/get-coupons")
			check(err)
	*/

	var coupons = Coupons{
		Coupon{
			"Tesco",
			2,
		},
		Coupon{
			"Sainsburys",
			2,
		},
	}

	couponsJson, err := json.Marshal(coupons)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(couponsJson)
}
