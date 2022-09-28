package pattern

import "fmt"

/*
	Реализовать паттерн «состояние». Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры
	использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// State is an interface that describe comp interface
type State interface {
	PressButton()
}

// comp is a struct contains comp values
type comp struct {
	currentState State
}

// NewComp is comp constructor
func NewComp() *comp {
	c := comp{}
	c.ChangeState(&turnedOff{cmp: &c})

	return &c
}

// PressButton is a function that executes PressButton function
func (c *comp) PressButton() {
	c.currentState.PressButton()
}

// ChangeState is a function that sets current comp state
func (c *comp) ChangeState(state State) {
	c.currentState = state
}

// turnedOff is a struct that contains comp
type turnedOff struct {
	cmp *comp
}

// PressButton is a function that executes PressButton function
func (t *turnedOff) PressButton() {
	fmt.Println("Nothing")

	t.cmp.ChangeState(&turnedOn{cmp: t.cmp})
}

// asleep is a struct that contains comp
type asleep struct {
	cmp *comp
}

// PressButton is a function that executes ChangeState function
func (a *asleep) PressButton() {
	fmt.Println("Wakes up the computer")

	a.cmp.ChangeState(&turnedOn{cmp: a.cmp})
}

// turnedOn is a struct that contains comp
type turnedOn struct {
	cmp *comp
}

// PressButton is a function that executes ChangeState function
func (t *turnedOn) PressButton() {
	fmt.Println("do something")

	t.cmp.ChangeState(&asleep{cmp: t.cmp})
}
