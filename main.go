package main

import "fmt"

func main() {
	err := fmt.Errorf("fdasfdsf")
	if nil != err {
		fmt.Println(err)
	}
}
