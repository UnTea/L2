package main

import (
	"strings"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testFindAnagrams := map[string][]string{
		"пятак":  {"пятак", "пятка", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
	}

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

	findAnagrams := FindAnagrams(dictionary)

	for key, value := range findAnagrams {
		var (
			anagrams []string
			ok       bool
		)

		if anagrams, ok = testFindAnagrams[key]; !ok {
			t.Errorf("excess key: %s", key)
		}

		joinedResultAnagrams := strings.Join(anagrams, " ")
		joinedMyAnagrams := strings.Join(value, " ")

		if joinedMyAnagrams != joinedResultAnagrams {
			t.Errorf("wrong anagrams:\nShould: %s\nGot: %s\n", joinedResultAnagrams, joinedMyAnagrams)
		}
	}
}
