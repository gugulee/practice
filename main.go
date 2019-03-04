package main

import (
	"fmt"
	"reflect"
)

func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})
	fmt.Println("hel")
	decoratedFunc.Set(v)
	return
}

type MyFoo func(int, int, int) int

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func main() {
	var myfoo MyFoo
	Decorator(&myfoo, foo)
	myfoo(1, 2, 3)
}
