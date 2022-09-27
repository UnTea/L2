package pattern

/*
	Реализовать паттерн «строитель». Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры
	использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type computer struct {
	CPU string
	RAM int
	MB  string
}

type ComputerBuilder interface {
	CPU(value string) ComputerBuilder
	RAM(value int) ComputerBuilder
	MB(value string) ComputerBuilder

	Build() computer
}

type computerBuilder struct {
	cpu string
	ram int
	mb  string
}

func NewComputerBuilder() ComputerBuilder {
	return computerBuilder{}
}

func (cb computerBuilder) CPU(value string) ComputerBuilder {
	cb.cpu = value

	return cb
}

func (cb computerBuilder) RAM(value int) ComputerBuilder {
	cb.ram = value

	return cb
}

func (cb computerBuilder) MB(value string) ComputerBuilder {
	cb.mb = value

	return cb
}

func (cb computerBuilder) Build() computer {
	return computer{
		CPU: cb.cpu,
		RAM: cb.ram,
		MB:  cb.mb,
	}
}

type serverComputerBuilder struct {
	computerBuilder
}

func NewServerComputerBuilder() ComputerBuilder {
	return serverComputerBuilder{}
}

func (scb serverComputerBuilder) Build() computer {
	return computer{
		CPU: "Xeon E5",
		RAM: 4,
		MB:  "GA-X99P-SLI (rev. 1.0)",
	}
}
