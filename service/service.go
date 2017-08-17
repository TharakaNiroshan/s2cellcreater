package service

import (
	//"compress/gzip"
	//"encoding/json"
	"fmt"
	"log"
	//"net/http"
	"strconv"

	"github.com/golang/geo/s2"
)

var (
	dataPoints  = []DataPoint{}
	heatMapData = []HeatMapRect{}
	cellMap     = map[s2.CellID]int{}
)

type CoordPair [2]float64
type Rect [4]CoordPair

type HeatMapRect struct {
	Rect      `json:"r"`
	Intensity int `json:"i"`
}

type DataPoint struct {
	lantitutte float64 `json:"lantitutte"`
	longtitute float64 `json:"longtitute"`
}

func IsCellidExisting(cellIdFromUser s2.CellID) bool {
	isExitCellId := false
	for cellId, _ := range cellMap {
		if cellId == cellIdFromUser {
			isExitCellId = true
		}
	}
	return isExitCellId
}

//GetSelectedPointCordinateAndGetSellID is a method which can use
//whether given cordinate in side the created polygon or not
//this method will return true if the given cordinate inside the created polygon
//if not will return false
func GetSelectedPointCordinateAndGetSellID(ars []string) bool {
	lantitute, err := strconv.ParseFloat(ars[0], 64)
	if err != nil {
		log.Fatal(err)
	}
	longtitute, err := strconv.ParseFloat(ars[1], 64)
	if err != nil {
		log.Fatal(err)
	}
	cellId := s2.CellIDFromLatLng(s2.LatLngFromDegrees(lantitute, longtitute)).Parent(18)
	boolean := IsCellidExisting(cellId)
	return boolean
}

//GetCodinate is the method getting the cordinate from user
//this will get the all the pointed cordinate from the string which is send by handler
//and parsing to CreateAreaByGivenCordinate method
//return []HeatMapRect
func GetHeatMap(requestData []string) []HeatMapRect {

	var getLength = len(requestData) / 2
	for i := 0; i < getLength; i++ {
		var k int = i * 2
		var l int = (1 + k)
		fmt.Println(requestData[k])
		fmt.Println(requestData[l])
		lantitute, err := strconv.ParseFloat(requestData[k], 64)
		if err != nil {
			log.Fatal(err)
		}
		longtitute, err := strconv.ParseFloat(requestData[l], 64)
		if err != nil {
			log.Fatal(err)
		}
		var dataPnt DataPoint
		dataPnt.lantitutte = lantitute
		dataPnt.longtitute = longtitute
		dataPoints = append(dataPoints, dataPnt)

	}
	var mapData = CreateAreaByGivenCordinate(dataPoints)
	return mapData
}

//CreateAreaByGivenCordinate is the one getting all the latitute,longtitute
//and creating a s2.Rect using those cordinates which is cover the given polygon
//and creating cells according to the created rect that  made by given cordinates in polygon

func CreateAreaByGivenCordinate(dataPoints []DataPoint) []HeatMapRect {
	const level = 18
	var rect s2.Rect
	for k := 0; k < len(dataPoints); k++ {
		if k == 0 {
			rect = s2.RectFromLatLng(s2.LatLngFromDegrees(dataPoints[k].lantitutte, dataPoints[k].longtitute))
		} else {
			rect = rect.AddPoint(s2.LatLngFromDegrees(dataPoints[k].lantitutte, dataPoints[k].longtitute))
		}
	}
	//creating region cover
	rc := &s2.RegionCoverer{MinLevel: level, MaxLevel: level, MaxCells: 50}
	//creating region accordinting to created rect
	r := s2.Region(rect.RectBound())
	//create cells covering area given coordinates
	covering := rc.Covering(r)

	for _, c := range covering {
		val := 8
		cellMap[c] += val
	}
	var heatMapRects = make([]HeatMapRect, 0, len(cellMap))
	for cellId, _ := range cellMap {
		//fmt.Println(val)
		var rect Rect
		cell := s2.CellFromCellID(cellId)
		for k := 0; k < 4; k++ {
			vertex := s2.LatLngFromPoint(cell.Vertex(k))
			rect[k][0], rect[k][1] = vertex.Lat.Degrees(), vertex.Lng.Degrees()
		}
		heatMapRects = append(heatMapRects, HeatMapRect{Rect: rect})
	}
	heatMapData = heatMapRects
	return heatMapData
}
