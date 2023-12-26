package main

import "fmt"

func main() {
	cond := map[string]interface{}{"id": "aa"}
	fmt.Println(cond)
	
	cond = map[string]interface{}{"ccc": "xxx"}
	fmt.Println(cond)
}