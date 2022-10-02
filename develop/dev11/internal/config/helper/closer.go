package helper

import (
	"io"
	"log"
)

// Closer is a function that closes connection with err handling
func Closer(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Println(err)
	}
}
