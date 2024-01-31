package test

import (
	"fmt"
	"nerversitup/services"
	"reflect"
	"strconv"
	"testing"
)

func TestOddInt(t *testing.T) {
	cases := [][]int{
		{0},
		{7},
		{1, 1, 2},
		{0, 1, 0, 1, 0},
		{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
	}
	for i, c := range cases {
		switch i {
		case 0:
			number, times := services.FindOddInt(c)
			expectedNumber, expectedTimes := 0, 1
			if !reflect.DeepEqual(number, expectedNumber) && !reflect.DeepEqual(times, expectedTimes) {
				t.Errorf("Case[%s][%v] error", strconv.Itoa(i), c)
			}
			t.Log("case 0 : success")
		case 1:
			number, times := services.FindOddInt(c)
			expectedNumber, expectedTimes := 7, 1
			if !reflect.DeepEqual(number, expectedNumber) && !reflect.DeepEqual(times, expectedTimes) {
				t.Errorf("Case[%s][%v] error", strconv.Itoa(i), c)
			}
			t.Log("case 0 : success")
		case 2:
			number, times := services.FindOddInt(c)
			expectedNumber, expectedTimes := 2, 1
			if !reflect.DeepEqual(number, expectedNumber) && !reflect.DeepEqual(times, expectedTimes) {
				t.Errorf("Case[%s][%v] error", strconv.Itoa(i), c)
			}
			t.Log("case 0 : success")
		case 3:
			number, times := services.FindOddInt(c)
			expectedNumber, expectedTimes := 1, 2
			if !reflect.DeepEqual(number, expectedNumber) && !reflect.DeepEqual(times, expectedTimes) {
				t.Errorf("Case[%s][%v] error", strconv.Itoa(i), c)
			}
			t.Log("case 0 : success")

		case 4:
			number, times := services.FindOddInt(c)
			expectedNumber, expectedTimes := 4, 1
			if !reflect.DeepEqual(number, expectedNumber) && !reflect.DeepEqual(times, expectedTimes) {
				t.Errorf("Case[%s][%v] error", strconv.Itoa(i), c)
			}
			t.Log("case 0 : success")
		}
	}
	fmt.Print(cases)
}
