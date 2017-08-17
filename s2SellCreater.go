package main

import (
	"log"
	"net/http"

	"pickme.lk/s2sellcreater/handler"
)

func main() {

	http.HandleFunc("/createpolygon", handler.GetPolygoneAndCreateArea)
	http.HandleFunc("/getselldetails", handler.GetSellDetails)
	// Start the server at http://localhost:8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
