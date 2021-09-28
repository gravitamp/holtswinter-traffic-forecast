package main

type Condition map[string]float64

type Classifier struct {
	datatrain []Condition
	datatest  []Condition
}

func NewClassifier() Classifier {
	c := Classifier{}
	return c
}

func (c *Classifier) addDataTrain(cond Condition) {
	c.datatrain = append(c.datatrain, cond)
}
func (c *Classifier) addDataTest(cond Condition) {
	c.datatest = append(c.datatest, cond)
}

func (c *Classifier) avg(param string, ctg float64) float64 {
	var total float64 = 0.0
	var jumlah float64 = 0.0
	for i := 0; i < len(c.datatrain); i++ {
		if c.datatrain[i]["days"] != ctg {
			continue
		}
		value, isFound := c.datatrain[i][param]
		if isFound {
			total += value
			jumlah += 1
		}
	}
	// fmt.Println("Jumlah : ", jumlah)

	return total
}
