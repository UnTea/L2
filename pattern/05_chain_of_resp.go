package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов». Объяснить применимость паттерна, его плюсы и минусы, а также реальные
	примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// status of problem
type problemStatus int

const (
	easy problemStatus = iota
	medium
	hard
	done
	unsolved
)

// problem is a function that contains status and description
type problem struct {
	status      problemStatus
	description string
}

// handler is an interface that defines list of methods
type handler interface {
	HandleProblem(*problem)
	SetNextHandler(handler)
}

// easyProblemHandler is a struct that handles all easy problems
type easyProblemHandler struct {
	next handler
}

// HandleProblem is a function that solves all easy problems
func (e *easyProblemHandler) HandleProblem(p *problem) {
	if p.status != easy {
		if e.next != nil {
			e.next.HandleProblem(p)
		} else {
			p.status = unsolved

			fmt.Printf("Cannot handle problem: %s", p.description)
		}
		return

	}

	p.status = done

	fmt.Printf("Problem: %s\nStatus: easy\nSolved\n", p.description)
}

// SetNextHandler is a function that sets next handler
func (e *easyProblemHandler) SetNextHandler(next handler) {
	e.next = next
}

// mediumProblemHandler is a struct that handles all medium problems
type mediumProblemHandler struct {
	next handler
}

// HandleProblem is a function that solves all medium problems
func (m *mediumProblemHandler) HandleProblem(p *problem) {
	if p.status != medium {
		if m.next != nil {
			m.next.HandleProblem(p)
		} else {
			p.status = unsolved

			fmt.Printf("Cannot handle problem: %s", p.description)
		}

		return
	}

	p.status = done

	fmt.Printf("Problem: %s\nStatus: medium\nSolved\n", p.description)
}

// SetNextHandler is a function that sets next handler
func (m *mediumProblemHandler) SetNextHandler(next handler) {
	m.next = next
}

// hardProblemHandler is a struct that handles all hard problems
type hardProblemHandler struct {
	next handler
}

// HandleProblem is a function that solves all hard problems
func (h *hardProblemHandler) HandleProblem(p *problem) {
	if p.status != hard {
		if h.next != nil {
			h.next.HandleProblem(p)
		} else {
			p.status = unsolved

			fmt.Printf("Cannot handle problem: %s", p.description)
		}

		return
	}

	p.status = done

	fmt.Printf("Problem: %s\nStatus: hard\nSolved\n", p.description)
}

// SetNextHandler is a function that sets next handler
func (h *hardProblemHandler) SetNextHandler(next handler) {
	h.next = next
}
