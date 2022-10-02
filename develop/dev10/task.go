package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

/*
	Реализовать простейший Telnet-клиент.

	Примеры вызовов:
	go-Telnet --timeout=10s host port go-Telnet mysite.ru 8080 go-Telnet --timeout=3s 1.1.1.1 123

	Требования:
	1. Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения
	STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
	2. Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию
	10s)
	3. При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера,
	программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через
	timeout
*/

type args struct {
	host    string
	port    string
	timeout time.Duration
}

func GetArgs() (*args, error) {
	if len(os.Args) < 3 {
		return nil, errors.New("error occurred: necessary to specify HOST and PORT")
	}

	var (
		host    string
		port    string
		timeout time.Duration
	)

	if strings.Contains(os.Args[1], "--timeout=") {
		timeSet := os.Args[1][len(os.Args[1])-1]
		if timeSet != 's' {
			return nil, errors.New("error occurred: necessary to specify time unit: e.g.: 10s")
		}

		index := strings.Index(os.Args[1], "=")
		num, err := strconv.Atoi(os.Args[1][index+1 : len(os.Args[1])-1])
		if err != nil || num < 1 {
			return nil, err
		}

		timeout, host, port = time.Duration(num)*time.Second, os.Args[2], os.Args[3]
	} else {
		timeout, host, port = time.Second*10, os.Args[1], os.Args[2]
	}

	return &args{
		host:    host,
		port:    port,
		timeout: timeout,
	}, nil
}

func ReadFromSocket(connect net.Conn, errChannel chan error) {
	input := make([]byte, 1024)

	for {
		n, err := connect.Read(input)
		if err != nil {
			errChannel <- fmt.Errorf("error occured: remote server stopped: %v", err)

			return
		}

		fmt.Println(string(input[:n]))
	}
}

func WriteToSocket(connection net.Conn, errChannel chan error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadBytes('\n')
		if err != nil {
			errChannel <- err

			return
		}

		text = text[:len(text)-1]

		_, err = connection.Write(text)
		if err != nil {
			errChannel <- err

			return
		}
	}
}

func Telnet(args *args) error {
	address := fmt.Sprintf("%s: %s", args.host, args.port)

	fmt.Println("Connecting to", address, "...")

	connection, err := net.DialTimeout("tcp", address, args.timeout)
	if err != nil {
		return err
	}

	defer func() {
		err = connection.Close()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("Connected to", address)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	errChannel := make(chan error)

	go ReadFromSocket(connection, errChannel)
	go WriteToSocket(connection, errChannel)

	select {
	case s := <-signals:
		fmt.Println("\nConnection stopped by signal:", s)
	case ec := <-errChannel:
		fmt.Println("Connection stopped by", ec)
	}

	return nil
}

func main() {
	parameters, err := GetArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = Telnet(parameters)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
