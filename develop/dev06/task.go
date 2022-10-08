package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	Утилита Cut

	Реализовать утилиту аналог консольной команды Cut (man Cut). Утилита должна принимать строки через STDIN, разбивать
	по разделителю (TAB) на колонки и выводить запрошенные.

	Реализовать поддержку утилитой следующих ключей:
	-f - "fields" - выбрать поля (колонки)
	-d - "delimiter" - использовать другой разделитель
	-s - "separated" - только строки с разделителем
*/

// Args is a function that contains arguments from CLI
type Args struct {
	f []int
	d string
	s bool

	files []string
}

// GetArgs is a function that returns *Args struct with parsed flags
func GetArgs() (*Args, error) {
	f := flag.String("f", "", "select only these fields; also print any line that contains no delimiter character, unless the -s option is specified")
	d := flag.String("d", "\t", "use DELIM instead of TAB for field delimite")
	s := flag.Bool("s", false, "do not print lines not containing delimiters")

	flag.Parse()

	if len(flag.Args()) < 1 {
		return nil, errors.New("you need to specify files")
	}

	if len(*f) < 1 {
		return nil, errors.New("you need to specify a field: e.g.: 1,3")
	}

	tmp := strings.Split(*f, ",")

	fields := make([]int, len(tmp))

	for i := range tmp {
		num, err := strconv.Atoi(tmp[i])
		if err != nil || num == 0 {
			return nil, fmt.Errorf("cannot convert string to int: %v", err)
		}
		fields[i] = num
	}

	args := &Args{
		f: fields,
		d: *d,
		s: *s,
	}

	args.files = append(args.files, flag.Args()...)

	return args, nil
}

// ReadFile is a function that reads file line by line
func ReadFile(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// CutLines is a function that cuts lines by specified args
func CutLines(args *Args, lines []string) []string {
	delimiter := "\t"

	if args.d != delimiter {
		delimiter = args.d
	}

	var result []string

	for _, line := range lines {
		if delimiter != "" && strings.Contains(line, delimiter) {
			words := strings.Split(line, delimiter)

			cutLine := strings.Builder{}

			for _, val := range args.f {
				if len(words) >= val {
					cutLine.WriteString(words[val-1])
					cutLine.WriteString(delimiter)
				}
			}

			result = append(result, strings.TrimSuffix(cutLine.String(), delimiter))

		} else if !args.s {
			result = append(result, line)
		}
	}

	return result
}

// Cut is a function that works like linux Cut with flags:
// -f -s -d
func Cut() ([]string, error) {
	if len(os.Args) < 4 {
		return nil, errors.New("not enough elements")
	}

	var result []string

	args, err := GetArgs()
	if err != nil {
		return nil, err
	}

	for _, val := range args.files {
		lines, err := ReadFile(val)
		if err != nil {
			return nil, err
		}

		cutLines := CutLines(args, lines)
		result = append(result, cutLines...)
	}

	return result, nil
}

func main() {
	lines, err := Cut()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i := range lines {
		fmt.Println(lines[i])
	}
}
