package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var vehiclestrain []float64
var vehiclestest []float64

func main() {
	setupData("trafficall1.csv")
	initialTrend(vehiclestrain, 24)
	initialSeasonalComponents(vehiclestrain, 168)
	forecast := TripleExponentialSmoothing(vehiclestrain, 168, 0.716, 0.029, 0.993, 24)
	acc := 0.0
	for i := 0; i < len(vehiclestest); i++ {
		fmt.Println("prediksi: ", forecast[i], "test :", vehiclestest[i])
		acc += (forecast[i] - vehiclestest[i])
	}
	fmt.Println("error:", acc/float64(len(vehiclestest)))
}

func setupData(file string) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	csvReader := csv.NewReader(f)
	csvData, err := csvReader.ReadAll()
	for i := 1; i < len(csvData); i++ {
		val, _ := strconv.ParseFloat(csvData[i][2], 64)
		//don't split randomly
		if i < len(csvData)-24 {
			vehiclestrain = append(vehiclestrain, val)
		} else {
			vehiclestest = append(vehiclestest, val)
		}
	}
}
