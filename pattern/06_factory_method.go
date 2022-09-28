package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод». Объяснить применимость паттерна, его плюсы и минусы, а также реальные
	примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Class - an abstract interface for mmo in game character class
type Class interface {
	GetClassName() string
}

// assassin is a struct that represents assassin character
type assassin struct {
	className     string
	characterName string
}

// NewAssassin is assassin constructor
func NewAssassin(name string) *assassin {
	return &assassin{
		className:     "Assassin",
		characterName: name,
	}
}

// GetClassName is a function that returns className field
func (a *assassin) GetClassName() string {
	return a.className
}

// warrior is a struct that represents a warrior character
type warrior struct {
	className     string
	characterName string
}

// NewWarrior is warrior constructor
func NewWarrior(name string) *warrior {
	return &warrior{
		className:     "Warrior",
		characterName: name,
	}
}

// GetClassName is a function that returns className field
func (w *warrior) GetClassName() string {
	return w.className
}

// GetClass is a fucntion that returns new MMO character
func GetClass(className, charName string) (Class, error) {
	switch className {
	case "Warrior":
		return NewWarrior(charName), nil
	case "Assassin":
		return NewAssassin(charName), nil
	}

	return nil, fmt.Errorf("wrong class name passed")
}
