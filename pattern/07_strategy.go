package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия». Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры
	использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// PowerOfGamer is an interface that describe character interface
type PowerOfGamer interface {
	Show(*character)
}

// character is a struct that contains character descriptions
type character struct {
	name         string
	lever        int
	strength     int
	agility      int
	intelligence int
	itemLevel    int
	power        PowerOfGamer
}

// NewGamer is character constructor
func NewGamer(name string, pow PowerOfGamer) *character {
	return &character{
		name:         name,
		lever:        1,
		strength:     10,
		agility:      12,
		intelligence: 8,
		itemLevel:    5,
		power:        pow,
	}
}

// SetPower is a function that sets character power
func (g *character) SetPower(p PowerOfGamer) {
	g.power = p
}

// ShowPower is a function that shows character power
func (g *character) ShowPower() {
	g.power.Show(g)
}

// powerByAverageStats is a struct that exists
type powerByAverageStats struct{}

// Show is a function that shows average character power by average stats
func (pbas *powerByAverageStats) Show(g *character) {
	pow := (g.agility + g.intelligence + g.strength) / 3.

	fmt.Printf("Power is equal to: %d\n", pow)
}

// powerByPlayerLevel is a struct that exists
type powerByPlayerLevel struct{}

// Show is a function that shows average character power by player lever
func (pbpl *powerByPlayerLevel) Show(g *character) {
	fmt.Printf("Power is equal to: %d\n", g.lever)
}

// powerByItemLever is a struct that exists
type powerByItemLever struct{}

// Show is a function that shows character power by item level
func (pbil *powerByItemLever) Show(g *character) {
	fmt.Printf("Power is equal to: %d\n", g.itemLevel)
}
