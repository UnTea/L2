package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
	Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

	- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
	- pwd - показать путь до текущего каталога
	- echo <args> - вывод аргумента в STDOUT
	- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
	- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


	Так же требуется поддерживать функционал fork/exec-команд

	Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

	*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое
	приглашение
	в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
	и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет
	введена команда выхода (например \quit).
*/

func command(strs ...string) {
	cmd := exec.Command(strs[0], strs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		cmd := scan.Text()

		args := strings.Split(cmd, " ")

		switch args[0] {
		case "cd":
			if len(args) == 1 {
				userDir, _ := os.UserHomeDir()
				os.Chdir(userDir)
			} else {
				os.Chdir(args[1])
			}
		default:
			if strings.ContainsRune(cmd, '|') {
				command("bash", "-c", cmd)
			} else {
				command(args...)
			}
		}
	}
}
