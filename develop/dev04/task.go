package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
	Поиск анаграмм по словарю

	Написать функцию поиска всех множеств анаграмм по словарю.


	Например:
	'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
	'листок', 'слиток' и 'столик' - другому.

	Требования:
	1. Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
	2. Выходные данные: ссылка на мапу множеств анаграмм
	3. Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
	слово из множества.
	4. Массив должен быть отсортирован по возрастанию.
	5. Множества из одного элемента не должны попасть в результат.
	6. Все слова должны быть приведены к нижнему регистру.
	7. В результате каждое слово должно встречаться только один раз.
*/

// FindAnagrams is a function that searching for anagrams in dictionary
func FindAnagrams(dictionary []string) map[string][]string {
	tmp := MakeTemporaryMapOfAnagrams(dictionary)

	Filter(tmp)

	anagrams := make(map[string][]string, len(tmp))

	for _, value := range tmp {
		sort.Strings(value)
		anagrams[value[0]] = value
	}

	return anagrams
}

// MakeTemporaryMapOfAnagrams is a function that makes temporary map of anagrams
func MakeTemporaryMapOfAnagrams(dictionary []string) map[string][]string {
	tmp := make(map[string][]string)

	for _, val := range dictionary {
		loweredWord := strings.ToLower(val)
		letters := GetLettersInAlphabetOrder(loweredWord)
		tmp[letters] = append(tmp[letters], loweredWord)
	}

	return tmp
}

// GetLettersInAlphabetOrder is a function that returns sorted in alphabet order string
func GetLettersInAlphabetOrder(word string) string {
	letters := strings.Split(word, "")
	sort.Strings(letters)

	return strings.Join(letters, "")
}

// Filter is a function that removes short sets and repeated words
func Filter(tmp map[string][]string) {
	unique := make(map[string]bool)

	for key, value := range tmp {
		// remove short sets
		if len(value) < 2 {
			delete(tmp, key)
		}

		// remove repeated words
		for i := range value {
			if !unique[value[i]] {
				unique[value[i]] = true
			} else {
				value[i] = value[len(value)-1]
				tmp[key] = value[:len(value)-1]
			}
		}

	}

}

func main() {
	dictionary := []string{
		"Пятак",
		"Пятак",
		"пятка",
		"Тяпка",
		"слиток",
		"слиток",
		"столик",
		"листок",
		"Топот",
		"Потоп",
	}

	anagrams := FindAnagrams(dictionary)

	for key, value := range anagrams {
		fmt.Printf("Key: %s\nValue: %v\n\n", key, value)
	}
}
