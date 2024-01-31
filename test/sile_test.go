package test

import (
	"fmt"
	"nerversitup/services"
	"reflect"
	"strconv"
	"testing"
)

func TestSmileCount(t *testing.T) {
	cases := [][]string{
		{":)", ";(", ";}", ":-D"},
		{";D", ":-(", ":-)", ";~D"},
		{":]", ":[", ";*", ";~D"},
	}
	for i, c := range cases {
		switch i {
		case 0:
			times := services.CountSmileys(c)
			expectedTimes := 2
			if !reflect.DeepEqual(times, expectedTimes) {
				t.Errorf("Case[%s][%v] error", strconv.Itoa(i), c)
			}
			t.Log("case 0 : success")
		case 1:
			times := services.CountSmileys(c)
			expectedTimes := 3
			if !reflect.DeepEqual(times, expectedTimes) {
				t.Errorf("Case[%s][%v] error", strconv.Itoa(i), c)
			}
			t.Log("case 1 : success")
		case 2:
			times := services.CountSmileys(c)
			expectedTimes := 1
			if !reflect.DeepEqual(times, expectedTimes) {
				t.Errorf("Case[%s][%v] error", strconv.Itoa(i), c)
			}
			t.Log("case 2 : success")
		}
	}
	fmt.Print(cases)
}
