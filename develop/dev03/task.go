package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
	Утилита sort
	Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные
	параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

	Реализовать поддержку утилитой следующих ключей:

	-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель
	— пробел)
	-n — сортировать по числовому значению
	-r — сортировать в обратном порядке
	-u — не выводить повторяющиеся строки

	Дополнительно

	Реализовать поддержку утилитой следующих ключей:

	-M — сортировать по названию месяца
	-b — игнорировать хвостовые пробелы
	-c — проверять отсортированы ли данные
	-h — сортировать по числовому значению с учетом суффиксов
*/

func usage() {
	fmt.Printf(`./sort [options] [filename]
	-k - specifying a column to sort
	-n - sort by numeric value
	-r - sort in reverse order
	-u - do not output duplicate lines
`)
}

func ReadFlags(columnFlag *int, numFlag, reverseFlag, uniqFlag *bool) {
	flag.IntVar(columnFlag, "k", 0, "specifying a column to sort")
	flag.BoolVar(numFlag, "n", false, "sort by numeric value")
	flag.BoolVar(reverseFlag, "r", false, "sort in reverse order")
	flag.BoolVar(uniqFlag, "u", false, "do not output duplicate lines")

	flag.Usage = usage

	flag.Parse()
}

func ReadFile() []string {
	args := flag.Args()

	if len(args) != 1 {
		flag.Usage()
		os.Exit(2)
	}

	filename := args[len(args)-1]

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	strs := strings.Split(string(bytes), "\n")

	return strs
}

func SortUtil(strs []string, columnFlag int, numFlag, reverseFlag, uniqFlag bool) []string {
	var uniqSet map[string]struct{}

	if uniqFlag {
		sliceSet := make([]string, 0)
		uniqSet = make(map[string]struct{})

		for _, str := range strs {
			_, ok := uniqSet[str]
			if !ok {
				sliceSet = append(sliceSet, str)
				uniqSet[str] = struct{}{}
			}
		}

		strs = sliceSet
	}

	if columnFlag != 0 {
		sort.Slice(strs, func(i, j int) bool {
			iStr := strings.Split(strs[i], " ")
			jStr := strings.Split(strs[j], " ")

			if len(iStr) < columnFlag && len(jStr) >= columnFlag {
				return true
			} else if len(jStr) < columnFlag && len(iStr) >= columnFlag {
				return false
			} else if len(iStr) < columnFlag && len(jStr) < columnFlag {
				return strs[i] < strs[j]
			}

			if numFlag {
				iNum, iErr := strconv.Atoi(iStr[columnFlag-1])
				jNum, jErr := strconv.Atoi(jStr[columnFlag-1])

				if iErr != nil || jErr != nil {
					return iStr[columnFlag-1] < jStr[columnFlag-1]
				}

				return iNum < jNum
			}

			return iStr[columnFlag-1] < jStr[columnFlag-1]
		})
	} else {
		sort.Strings(strs)
	}

	if reverseFlag {
		for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
			strs[i], strs[j] = strs[j], strs[i]
		}
	}

	return strs
}

func main() {
	var (
		columnFlag  int
		numFlag     bool
		reverseFlag bool
		uniqFlag    bool
	)

	ReadFlags(&columnFlag, &numFlag, &reverseFlag, &uniqFlag)

	strs := ReadFile()
	strs = SortUtil(strs, columnFlag, numFlag, reverseFlag, uniqFlag)

	for _, str := range strs {
		fmt.Println(str)
	}
}
