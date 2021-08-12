package naive_bayesian

import (
	"fmt"
	"testing"
)

var appleSample = []fruit{
	{
		shape:  1,
		color:  1,
		grain:  1,
		weight: 2,
		grip:   1,
		taste:  1,
	},
	{
		shape:  1,
		color:  1,
		grain:  1,
		weight: 1,
		grip:   1,
		taste:  1,
	},
	{
		shape:  2,
		color:  3,
		grain:  1,
		weight: 1,
		grip:   2,
		taste:  1,
	},
}

var orangeSample = []fruit{
	{
		shape:  2,
		color:  2,
		grain:  1,
		weight: 1,
		grip:   2,
		taste:  2,
	},
	{
		shape:  2,
		color:  2,
		grain:  1,
		weight: 2,
		grip:   2,
		taste:  2,
	},
	{
		shape:  1,
		color:  2,
		grain:  1,
		weight: 2,
		grip:   1,
		taste:  1,
	},
}

var melonSample = []fruit{
	{
		shape:  3,
		color:  3,
		grain:  2,
		weight: 3,
		grip:   1,
		taste:  2,
	},
	{
		shape:  3,
		color:  3,
		grain:  2,
		weight: 3,
		grip:   1,
		taste:  1,
	},
	{
		shape:  3,
		color:  3,
		grain:  2,
		weight: 3,
		grip:   1,
		taste:  2,
	},
	{
		shape:  1,
		color:  3,
		grain:  2,
		weight: 3,
		grip:   2,
		taste:  2,
	},
}

func TestAttributeAtClass(t *testing.T) {
	pAttrAtApple := AttributeAtClass(appleSample)
	pAttrAtOrange := AttributeAtClass(orangeSample)
	pAttrAtMelon := AttributeAtClass(melonSample)

	pApple := ClassInTotal(appleSample, orangeSample, melonSample)
	pOrange := ClassInTotal(orangeSample, appleSample, melonSample)
	pMelon := ClassInTotal(melonSample, appleSample, orangeSample)

	pAttr := AttributeAtTotal(appleSample, orangeSample, melonSample)

	pIsApple := ClassAtAttribute(map[string]string{"shape": "round", "taste": "sweet"}, pApple, pAttrAtApple, pAttr)
	fmt.Printf("the probability of apple is: %f\n", pIsApple)

	pIsOrange := ClassAtAttribute(map[string]string{"shape": "round", "taste": "sweet"}, pOrange, pAttrAtOrange, pAttr)
	fmt.Printf("the probability of orange is: %f\n", pIsOrange)

	pIsMelon := ClassAtAttribute(map[string]string{"shape": "round", "taste": "sweet"}, pMelon, pAttrAtMelon, pAttr)
	fmt.Printf("the probability of melon is: %f\n", pIsMelon)
}
