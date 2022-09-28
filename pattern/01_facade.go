package pattern

import (
	"errors"
	"math/rand"
	"time"
)

/*
	Реализовать паттерн «фасад». Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры
	использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// email is a struct that represents email
type email struct {
	login  string
	domain string
}

// NewEmail is a constructor of email structure
func NewEmail(l, d string) *email {
	return &email{
		login:  l,
		domain: d,
	}
}

// Check is a function that returns result checking the correctness of the incoming login and domain
func (e *email) Check(incomingLogin, incomingDomain string) (result bool) {
	if e.login != incomingLogin || e.domain != incomingDomain {
		return result
	}

	return !result
}

// password is a struct that represents password
type password struct {
	password string
}

// NewPassword is a constructor of password structure
func NewPassword(p string) *password {
	return &password{
		password: p,
	}
}

// Check is a function that returns result of checking the correctness of the incoming password
func (p *password) Check(incomingPassword string) (result bool) {
	if p.password != incomingPassword {
		return
	}

	return !result
}

// securityCode is a struct that represents securityCode
type securityCode struct {
	sc int
}

// NewSecurityCode is a constructor of securityCode that creates new securityCode instance and sets his default value
func NewSecurityCode() *securityCode {
	return &securityCode{
		sc: 0,
	}
}

// Check is a function that returns result of checking the correctness of the incoming security code
func (s *securityCode) Check(incomingCode int) (result bool) {
	if s.sc != incomingCode {
		return
	}

	return !result
}

// SendCode is a function that returns random generated code for security check
func (s *securityCode) SendCode() (code int) {
	rand.Seed(time.Now().UnixNano())
	code = rand.Int()
	s.sc = code

	return
}

// user is a struct that represents user
type user struct {
	email        *email
	password     *password
	securityCode *securityCode
}

// NewUser is a constructor of user structure
func NewUser(login, domain, password string) *user {
	return &user{
		email:        NewEmail(login, domain),
		password:     NewPassword(password),
		securityCode: NewSecurityCode(),
	}
}

// Login is a function that hides all internal processes and gives user only required interface
func (u *user) Login(login, domain, password string) error {
	if u.email.Check(login, domain) && u.password.Check(password) {
		code := u.securityCode.SendCode()

		if !u.securityCode.Check(code) {
			return errors.New("error occurred while checking security code")
		}

		return nil
	}

	return errors.New("error occurred while checking input values")
}
