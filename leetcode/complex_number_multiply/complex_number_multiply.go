package cnm

import "fmt"

// 	num1 := "a+bi"
//	num2 := "c+di"
// (a+bi) *(c+di)
func complexNumberMultiply(num1 string, num2 string) string {
	var a, b, c, d int
	fmt.Sscanf(num1, "%d+%di", &a, &b)
	fmt.Sscanf(num2, "%d+%di", &c, &d)
	fmt.Println(a, b, c, d)

	return fmt.Sprintf("%d+%di", a*c-b*d, a*d+b*c)
}
