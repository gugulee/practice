package naive_bayesian

import "math"

var (
	shape  = map[string]int{"irregular": 1, "round": 2, "oval": 3}
	color  = map[string]int{"red": 1, "orange": 2, "green": 3}
	grain  = map[string]int{"noGrain": 1, "grain": 2}
	weight = map[string]int{"0-200": 1, "200-500": 2, "500>1000": 3}
	grip   = map[string]int{"hard": 1, "soft": 2}
	taste  = map[string]int{"sour": 1, "sweet": 2}
)

type fruit struct {
	shape  int
	color  int
	grain  int
	weight int
	grip   int
	taste  int
}

type Probability map[string]PAttr

type PAttr map[string]float64

// classInTotal calculate the probability of of P(Class)
// total[0] is the target which calculate probability
func ClassInTotal(total ...[]fruit) float64 {
	class := total[0]
	classCount := len(class)
	var totalCount int
	for i := range total {
		for range total[i] {
			totalCount++
		}
	}
	return float64(classCount) / float64(totalCount)
}

// AttributeAtTotal calculate the probability of of P(Attribute)
func AttributeAtTotal(total ...[]fruit) *Probability {
	var totalCount int
	for i := range total {
		for range total[i] {
			totalCount++
		}
	}

	var shapeCount = map[int]int{}
	var colorCount = map[int]int{}
	var grainCount = map[int]int{}
	var weightCount = map[int]int{}
	var gripCount = map[int]int{}
	var tasteCount = map[int]int{}
	var probability = Probability{
		"shape":  PAttr{},
		"color":  PAttr{},
		"grain":  PAttr{},
		"weight": PAttr{},
		"grip":   PAttr{},
		"taste":  PAttr{},
	}

	for _, class := range total {
		for j := range class {
			shapeCount[class[j].shape]++
			colorCount[class[j].color]++
			grainCount[class[j].grain]++
			weightCount[class[j].weight]++
			gripCount[class[j].grip]++
			tasteCount[class[j].taste]++
		}
	}

	calculateProbability("shape", shapeCount, shape, &probability, totalCount)
	calculateProbability("color", colorCount, color, &probability, totalCount)
	calculateProbability("grain", grainCount, grain, &probability, totalCount)
	calculateProbability("weight", weightCount, weight, &probability, totalCount)
	calculateProbability("grip", gripCount, grip, &probability, totalCount)
	calculateProbability("taste", tasteCount, taste, &probability, totalCount)

	return &probability
}

// ClassAtAttribute calculate the probability of P(Class|Attribute)
// property["shape"]=irregular
func ClassAtAttribute(properties map[string]string, pClass float64, pAttrAtClass, pAttrAtTotal *Probability) float64 {
	var p float64 = 1
	for property, attr := range properties {
		p1 := (*pAttrAtClass)[property][attr]
		p2 := (*pAttrAtTotal)[property][attr]
		if isEqual(p1, 0) {
			p1 = 0.01
		}

		if isEqual(p2, 0) {
			p2 = 0.01
		}

		p *= p1 * pClass / p2
	}

	return p
}

// AttributeAtClass calculate the probability of P(Attribute|Class)
func AttributeAtClass(class []fruit) *Probability {
	fruitCount := len(class)
	var shapeCount = map[int]int{}
	var colorCount = map[int]int{}
	var grainCount = map[int]int{}
	var weightCount = map[int]int{}
	var gripCount = map[int]int{}
	var tasteCount = map[int]int{}
	var probability = Probability{
		"shape":  PAttr{},
		"color":  PAttr{},
		"grain":  PAttr{},
		"weight": PAttr{},
		"grip":   PAttr{},
		"taste":  PAttr{},
	}

	for _, f := range class {
		shapeCount[f.shape]++
		colorCount[f.color]++
		grainCount[f.grain]++
		weightCount[f.weight]++
		gripCount[f.grip]++
		tasteCount[f.taste]++
	}

	calculateProbability("shape", shapeCount, shape, &probability, fruitCount)
	calculateProbability("color", colorCount, color, &probability, fruitCount)
	calculateProbability("grain", grainCount, grain, &probability, fruitCount)
	calculateProbability("weight", weightCount, weight, &probability, fruitCount)
	calculateProbability("grip", gripCount, grip, &probability, fruitCount)
	calculateProbability("taste", tasteCount, taste, &probability, fruitCount)

	return &probability
}

func calculateProbability(property string, count map[int]int, relations map[string]int, p *Probability, total int) {
	for mark1, count := range count {
		for attr, mark2 := range relations {
			if mark1 == mark2 {
				(*p)[property][attr] = float64(count) / float64(total)
			}
		}
	}

}

func isEqual(f1, f2 float64) bool {
	return math.Dim(f1, f2) < 0.0000001
}
