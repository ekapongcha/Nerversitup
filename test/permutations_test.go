package test

import (
	"fmt"
	"nerversitup/services"
	"reflect"
	"strconv"
	"testing"
)

func TestPermutation(t *testing.T) {
	cases := []string{"A", "AB", "ABC"}
	for i, c := range cases {
		switch i {
		case 0:
			input := services.Permutation(c)
			expected := []string{"A"}
			if !reflect.DeepEqual(input, expected) {
				t.Errorf("Case[%s][%s] error", strconv.Itoa(i), c)
			}
			t.Log("case 0 : success")
		case 1:
			input := services.Permutation(c)
			expected := []string{"AB", "BA"}
			if !reflect.DeepEqual(input, expected) {
				t.Errorf("Case[%s][%s] error", strconv.Itoa(i), c)
			}
			t.Log("case 1 : success")
		case 2:
			input := services.Permutation(c)
			expected := []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"}
			if !reflect.DeepEqual(input, expected) {
				t.Errorf("Case[%s][%s] error", strconv.Itoa(i), c)
			}
			t.Log("case 2 : success")
		}
	}
	fmt.Print(cases)
}
