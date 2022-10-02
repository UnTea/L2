package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func HandleConnection(ctx context.Context, connection net.Conn) {
	defer func() {
		err := connection.Close()
		if err != nil {
			panic(err)
		}
	}()

	input := make([]byte, 1024)

	for {
		n, err := connection.Read(input)
		if err == io.EOF {
			break
		} else if n == 0 || err != nil {
			log.Fatal(err)
		}

		_, err = connection.Write(append([]byte("Server: you wrote: "), input[:n]...))
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Connection with %s closed\n", connection.RemoteAddr())
}

const address = "localhost:8080"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Starting server ...")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = listener.Close()
		if err != nil {
			panic(err)
		}
	}()

	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-signals
		cancel()

		fmt.Println("\nServer stopped by signal: ", s)
		os.Exit(0)
	}()

	for {
		connection, err := listener.Accept()
		if err != nil {
			err := connection.Close()
			if err != nil {
				panic(err)
			}

			log.Fatal(err)
		}

		fmt.Printf("Connection with %s established\n", connection.RemoteAddr())

		go HandleConnection(ctx, connection)
	}
}
