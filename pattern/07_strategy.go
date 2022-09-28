package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия». Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры
	использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type PowerOfGamer interface {
	Show(*gamer)
}

type gamer struct {
	name         string
	lever        int
	strength     int
	agility      int
	intelligence int
	itemLevel    int
	power        PowerOfGamer
}

func NewGamer(name string, pow PowerOfGamer) *gamer {
	return &gamer{
		name:         name,
		lever:        1,
		strength:     10,
		agility:      20,
		intelligence: 50,
		itemLevel:    10,
		power:        pow,
	}
}

func (g *gamer) SetPower(p PowerOfGamer) {
	g.power = p
}

func (g *gamer) ShowPower() {
	g.power.Show(g)
}

type powerByAverageStats struct{}

func (pbas *powerByAverageStats) Show(g *gamer) {
	pow := (g.agility + g.intelligence + g.strength) / 3.

	fmt.Printf("Power is equal to: %d\n", pow)
}

type powerByPlayerLevel struct{}

func (pbpl *powerByPlayerLevel) Show(g *gamer) {
	fmt.Printf("Power is equal to: %d\n", g.lever)
}

type powerByItemLever struct{}

func (pbil *powerByItemLever) Show(g *gamer) {
	fmt.Printf("Power is equal to: %d\n", g.itemLevel)
}
