package handler

import (
	"compress/gzip"
	"encoding/json"
	//	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"pickme.lk/s2sellcreater/service"
)

type CoordPair [2]float64
type Rect [4]CoordPair
type HeatMapRect struct {
	Rect      `json:"r"`
	Intensity int `json:"i"`
}

func GetPolygoneAndCreateArea(w http.ResponseWriter, r *http.Request) {
	req, _ := ioutil.ReadAll(r.Body)
	var s string = strings.Replace(string(req), "[", "", 1)
	s = strings.Replace(s, "]", "", 1)
	s = strings.Replace(s, "\"", "", -1)
	s = strings.Replace(s, " ", "", -1)
	var ars []string = strings.Split(s, ",")
	var heatMapRect []service.HeatMapRect = service.GetHeatMap(ars)
	sendResponseJSON(w, heatMapRect)
}

func GetSellDetails(w http.ResponseWriter, r *http.Request) {
	req, _ := ioutil.ReadAll(r.Body)
	dataString := string(req)
	var s string = strings.Replace(dataString, "\"", "", -1)
	var ars []string = strings.Split(s, ",")
	boolean := service.GetSelectedPointCordinateAndGetSellID(ars)
	log.Println("if point inside the polygon : ", boolean)

}

func sendResponseJSON(res http.ResponseWriter, resData interface{}) {
	res.Header().Add("Access-Control-Allow-Origin", "*")
	res.Header().Add("Content-Type", "application/json")
	res.Header().Set("Content-Encoding", "gzip")
	data, _ := json.Marshal(resData)
	writer := gzip.NewWriter(res)
	defer writer.Close()
	log.Println(data)
	writer.Write(data)
}
