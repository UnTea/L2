package main

import (
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	passingCase := time.Now()

	failureCases := []time.Time{
		time.Now().Add(time.Hour * 2),
		time.Now().Add(time.Minute * 2),
		time.Now().Add(time.Second * 2),
	}

	currentTime, err := GetCurrentTime()
	if err != nil {
		t.Error(err)
	}

	if currentTime.Before(passingCase) {
		t.Errorf("%v != %v", passingCase, currentTime)
	}

	for _, value := range failureCases {
		if value.Before(currentTime) {
			t.Errorf("%v == %v\nShould: %v\nGot: %v\n", currentTime, value, currentTime, value)
		}
	}
}
