package template

import (
	"fmt"
	"time"
)

type Runner struct {
	doRun func()
}

func (r *Runner) Start() {
	r.doRun()
}

type Logger struct {
	Runner
}

func NewLogger() *Logger {
	l := &Logger{}
	l.doRun = l.doRunImpl
	return l
}

func (l *Logger) doRunImpl() {
	time.Sleep(1 * time.Second)
	fmt.Println("Running")
}