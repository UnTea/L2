package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель». Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры
	использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// RandomEvents is an interface that holds all methods that will be applied on all live forms
type RandomEvents interface {
	EventForPlayer(p *player) string
	EventForGoblin(g *goblin) string
}

// LiveForms is an interface for accepting new methods without changing struct
type LiveForms interface {
	Accept(event RandomEvents)
}

// player is a struct that holds player's data
type player struct {
	nickname string
	hp       int
	mana     int
	exp      int
}

// NewPlayer ia a player construct
func NewPlayer(nickname string) *player {
	return &player{
		nickname: nickname,
		hp:       100,
		mana:     100,
		exp:      0,
	}
}

// GetDamage is a function that applies damage on player
func (p *player) GetDamage(dmg int) {
	p.hp -= dmg

	fmt.Printf("player %s was attacked on %d and now have %d health\n", p.nickname, dmg, p.hp)
}

// Accept is a function that applies random effect on a player
func (p *player) Accept(event RandomEvents) {
	event.EventForPlayer(p)
}

// goblin is a struct that contains goblin's data
type goblin struct {
	damage int
	hp     int
	lvl    int
}

// NewGoblin ia a goblin construct
func NewGoblin(damage int) *goblin {
	return &goblin{
		damage: damage,
		hp:     100,
		lvl:    1,
	}
}

// Attack is a function that deals damage to player
func (g *goblin) Attack(p *player) {
	fmt.Printf("goblin has attacked %s on damage %d\n", p.nickname, g.damage)
}

// Accept is a function that applies random effect on a goblin
func (g *goblin) Accept(event RandomEvents) {
	event.EventForGoblin(g)
}

// (visitor) holyPotato is a struct that emits holy light that heals human wounds and damages goblins
type holyPotato struct {
	damage int
	heal   int
}

// NewHolyPotato is a holyPotato construct
func NewHolyPotato(dmg, heal int) *holyPotato {
	return &holyPotato{
		damage: dmg,
		heal:   heal,
	}
}

// EventForPlayer is a function that applies event on player
func (h *holyPotato) EventForPlayer(p *player) string {
	p.hp += h.heal
	return fmt.Sprintf("Holy potato heals %s on %d healp points", p.nickname, h.heal)
}

// EventForGoblin is a function that applies event on goblin
func (h *holyPotato) EventForGoblin(g *goblin) string {
	g.hp -= h.damage
	return fmt.Sprintf("Holy potato damages goblin on %d points", h.damage)
}
