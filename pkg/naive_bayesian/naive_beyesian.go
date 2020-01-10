package naive_bayesian

var (
	shape  = map[string]int{"irregular": 1, "round": 2, "oval": 3}
	color  = map[string]int{"red": 1, "orange": 2, "green": 3}
	grain  = map[string]int{"noGrain": 1, "grain": 2}
	weight = map[string]int{"0-200": 1, "200-500": 2, "500>1000": 3}
	grip   = map[string]int{"hard": 1, "soft": 2}
	taste  = map[string]int{"sour": 1, "sweet": 2}
)

// AttributeAtClass calculate the probability of P(Attribute|Class)
func AttributeAtClass(attr, class string) float32 {

	return 0
}
