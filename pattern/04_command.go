package pattern

/*
	Реализовать паттерн «комманда». Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры
	использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Command is an interface
type Command interface {
	Execute()
}

// Seller is an interface
type Seller interface {
	Sell(amount int)
	Buy(amount int)
}

// sellCommand is a struct that contains seller interface and amount
type sellCommand struct {
	seller Seller
	amount int
}

// NewSellCommand is a function that creates new instance with all needed parameters
func NewSellCommand(seller Seller, amount int) *sellCommand {
	return &sellCommand{
		seller: seller,
		amount: amount,
	}
}

// Execute is a function that
func (s *sellCommand) Execute() {
	s.seller.Sell(s.amount)
}

// buyCommand is a struct that contains seller interface and amount
type buyCommand struct {
	seller Seller
	amount int
}

// NewBuyCommand - create new instance with all needed params
func NewBuyCommand(seller Seller, amount int) *buyCommand {
	return &buyCommand{
		amount: amount,
		seller: seller,
	}
}

// Execute is a function that executes given command
func (b *buyCommand) Execute() {
	b.seller.Buy(b.amount)
}

// bot is a struct that represent sender
type bot struct {
	cmd Command
}

// NewBot is a function that creates new instance of Bot
func NewBot(cmd Command) *bot {
	return &bot{
		cmd: cmd,
	}
}

// SetCommand is a function that sets new command in bot
func (b *bot) SetCommand(cmd Command) {
	b.cmd = cmd
}

// Start is a function that starts execution function
func (b *bot) Start() {
	b.cmd.Execute()
}

// ivan is a struct that represent receiver
type ivan struct {
	items int
}

// Sell is a function that decrement amount of items
func (i *ivan) Sell(amount int) {
	i.items -= amount
}

// Buy is a function that increase amount of items
func (i *ivan) Buy(amount int) {
	i.items += amount
}
