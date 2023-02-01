package v2

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Visitor interface {
	visit(Shape)
}

type Shape interface {
	accept(Visitor)
}

type Circle struct {
	Radius int
}

func (c Circle) accept(v Visitor) {
	v.visit(c)
}

type Rectangle struct {
	Width, Heigh int
}

func (r Rectangle) accept(v Visitor) {
	v.visit(r)
}

type JsonVisitor struct{}

func (v JsonVisitor) visit(shape Shape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

type XmlVisitor struct{}

func (v XmlVisitor) visit(shape Shape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
