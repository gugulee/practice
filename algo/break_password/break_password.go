package break_password

import "fmt"

var letters = "abcdefghijklmnopqrstuvwxyz"

func breakPassword(password, result string) bool {
	if len(password) == len(result) {
		if password == result {
			fmt.Println(result)
			return true
		}
		return false
	}

	for _, letter := range letters {
		netResult := result
		netResult += string(letter)
		if found := breakPassword(password, netResult); found {
			return true
		}
	}
	return false
}

//func breakPassword(password int, result string) {
//	if 0 == password {
//		a++
//		fmt.Println(a, result)
//		return
//	}
//
//	for _, letter := range letters {
//		netResult := result
//		netResult += string(letter)
//		breakPassword(password-1, netResult)
//	}
//}
