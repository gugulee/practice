package template

import "fmt"

type Game interface {
	initialize()
	start()
	end()
}

type Play struct {
	Game
}

func New(g Game) *Play {
	return &Play{Game: g}
}

func (p *Play) Play() {
	p.initialize()
	p.start()
	p.end()
}

type Cricket struct{}

func (c *Cricket) initialize() {
	fmt.Println("cricket initialize")
}

func (c *Cricket) start() {
	fmt.Println("cricket start")
}

func (c *Cricket) end() {
	fmt.Println("cricket end")
}

type Football struct{}

func (f *Football) initialize() {
	fmt.Println("football initialize")
}

func (f *Football) start() {
	fmt.Println("football start")
}

func (f *Football) end() {
	fmt.Println("football end")
}
