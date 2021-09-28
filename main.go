package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	MONDAY    = 1.0
	TUESDAY   = 2.0
	WEDNESDAY = 3.0
	THURSDAY  = 4.0
	FRIDAY    = 5.0
	SATURDAY  = 6.0
	SUNDAY    = 7.0
)

var vehiclestrain []float64
var vehiclestest []float64

func main() {
	c := NewClassifier()
	c.setupData("trafficall1.csv")
	initialTrend(vehiclestrain, 24)
	initialSeasonalComponents(vehiclestrain, 168)
	forecast := TripleExponentialSmoothing(vehiclestrain, 168, 0.716, 0.029, 0.993, 24)
	acc := 0.0
	for i := 0; i < len(vehiclestest); i++ {
		fmt.Println("prediksi: ", forecast[i], "test :", vehiclestest[i])
		acc += math.Abs(forecast[i] - vehiclestest[i])
	}
	fmt.Println("prediksi jam 8: ", forecast[7])
	fmt.Println("error:", acc/float64(len(vehiclestest)))
	fmt.Println("comparing")
	fmt.Println("Sum monday : ", c.avg("vehicles", MONDAY))
	fmt.Println("Sum tuesday : ", c.avg("vehicles", TUESDAY))
	fmt.Println("Sum wednesday : ", c.avg("vehicles", WEDNESDAY))
	fmt.Println("Sum thursday : ", c.avg("vehicles", THURSDAY))
	fmt.Println("Sum friday : ", c.avg("vehicles", FRIDAY))
	fmt.Println("Sum saturday : ", c.avg("vehicles", SATURDAY))
	fmt.Println("Sum monday : ", c.avg("vehicles", SUNDAY))
}

func (c *Classifier) setupData(file string) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	csvReader := csv.NewReader(f)
	csvData, err := csvReader.ReadAll()
	for i := 1; i < len(csvData); i++ {
		days := MONDAY
		switch csvData[i][4] {
		case "Monday":
			days = MONDAY
			break
		case "Tuesday":
			days = TUESDAY
			break
		case "Wednesday":
			days = WEDNESDAY
			break
		case "Thursday":
			days = THURSDAY
			break
		case "Friday":
			days = FRIDAY
			break
		case "Saturday":
			days = SATURDAY
			break
		case "Sunday":
			days = SUNDAY
			break
		}
		val, _ := strconv.ParseFloat(csvData[i][2], 64)
		//don't split randomly
		if i < len(csvData)-24 {
			vehiclestrain = append(vehiclestrain, val)
			c.addDataTrain(Condition{
				"vehicles": val,
				"days":     days,
			})
		} else {
			vehiclestest = append(vehiclestest, val)
			c.addDataTest(Condition{
				"vehicles": val,
				"days":     days,
			})
		}
	}
}
