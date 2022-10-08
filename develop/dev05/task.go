package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
	Реализовать утилиту фильтрации (man grep)
	Поддержать флаги:
	-A - "after" печатать +N строк после совпадения
	-B - "before" печатать +N строк до совпадения
	-C - "context" (A+B) печатать ±N строк вокруг совпадения
	-c - "count" (количество строк)
	-i - "ignore-case" (игнорировать регистр)
	-v - "invert" (вместо совпадения, исключать)
	-F - "fixed", точное совпадение со строкой, не паттерн
	-n - "line num", печатать номер строки
*/

type customGrep struct {
	flags    customFlags
	regExp   string
	filename string
	content  []fileContent
}

type customFlags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	index      bool
}

type fileContent struct {
	text  string
	index int
}

func Usage() {
	fmt.Println("./grep [flags] pattern filename")
}

func (g *customGrep) ReadFlags() {
	flgs := customFlags{}
	flag.IntVar(&flgs.after, "A", 0, "'after' print +N lines after match")
	flag.IntVar(&flgs.before, "B", 0, "'before' print +N lines until match")
	flag.IntVar(&flgs.context, "C", 0, "'context' (A+B) print ±N lines around the match")
	flag.BoolVar(&flgs.count, "c", false, "'count' (number of lines)")
	flag.BoolVar(&flgs.ignoreCase, "i", false, "'ignore-case' (ignore case)")
	flag.BoolVar(&flgs.invert, "v", false, "'invert' (instead of match, exclude)")
	flag.BoolVar(&flgs.fixed, "F", false, "'fixed', exact string match, not a pattern")
	flag.BoolVar(&flgs.index, "n", false, "'line num', print line number")

	flag.Usage = Usage

	flag.Parse()

	args := flag.Args()
	if flgs.context > 0 {
		flgs.after, flgs.before = flgs.context, flgs.context
	}

	if flgs.count {
		flgs.after, flgs.before = 0, 0
	}

	if len(args) == 2 {
		g.regExp = args[0]
		g.filename = args[1]
	} else {
		Usage()
		os.Exit(0)
	}

	g.flags = flgs
}

func ReadFile(filename string) ([]fileContent, error) {
	var rows []fileContent
	file, err := os.Open(filename)
	if err != nil {
		return rows, err
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	count := 1

	for sc.Scan() {
		rows = append(rows, fileContent{
			text:  sc.Text(),
			index: count,
		})
		count++
	}

	return rows, nil
}

func match(fixed bool, reg *regexp.Regexp, str string) bool {
	if fixed {
		return reg.String() == str
	}

	return reg.MatchString(str)
}

func (g *customGrep) Grep() (string, error) {
	var prefix, postfix string

	if g.flags.ignoreCase {
		prefix = "(?i)"
	}

	re, err := regexp.Compile(prefix + g.regExp + postfix)

	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	indexStr := make([]int, 0)

	if g.flags.invert {
		for i, str := range g.content {
			if !re.MatchString(str.text) {
				indexStr = append(indexStr, i)
			}
		}
	} else {
		for i, str := range g.content {
			if re.MatchString(str.text) {
				indexStr = append(indexStr, i)
			}
		}
	}

	if len(indexStr) == 0 {
		return "", nil
	}

	result := make([]fileContent, 0, len(indexStr))

	switch {
	case g.flags.after > 0 || g.flags.before > 0:
		after, before := g.flags.after, g.flags.before
		start, end := indexStr[0]-before, indexStr[0]+after

		if indexStr[0]-before < 0 {
			start = 0
		}

		if end >= len(g.content) {
			end = len(g.content) - 1
		}

		result = append(result, g.content[start:end+1]...)

		if len(indexStr) > 1 {
			for _, index := range indexStr[1:] {

				if before != 0 {

					if index > end && end >= index-before {
						result = append(result, g.content[end+1:index+1]...)
						end = index
					}
					if index > end && end < index-before {
						result = append(result, g.content[index-before:index+1]...)
						end = index
					}
				}

				if index > end {
					result = append(result, g.content[index])
				}

				if after != 0 {
					lastA := end
					if len(g.content) <= index+after {
						end = len(g.content) - 1
					} else {
						end = index + after
					}
					if index <= lastA && lastA <= index+after {
						result = append(result, g.content[lastA+1:end+1]...)
						continue
					}
					result = append(result, g.content[index+1:end+1]...)
				}

			}

		}
	case g.flags.count:
		return strconv.Itoa(len(indexStr)), nil
	default:
		for _, index := range indexStr {
			result = append(result, g.content[index])
		}
	}

	return Join(result, g.flags.index), nil
}

func Join(elems []fileContent, strNum bool) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		if strNum {
			return strconv.Itoa(elems[0].index) + ": " + elems[0].text
		}

		return elems[0].text
	}

	sep := "\n"

	var b strings.Builder

	if strNum {
		b.WriteString(strconv.Itoa(elems[0].index) + ": " + elems[0].text)

		for _, s := range elems[1:] {
			b.WriteString(sep)
			b.WriteString(strconv.Itoa(s.index) + ": " + s.text)
		}

		return b.String()
	}

	b.WriteString(elems[0].text)

	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s.text)
	}

	return b.String()
}

func main() {
	grep := &customGrep{}
	lines, err := ReadFile("text.txt")

	if err != nil {
		log.Fatalln(err)
	}

	grep.ReadFlags()

	grep.content = lines
	res, err := grep.Grep()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}
