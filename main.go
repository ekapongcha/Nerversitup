package main

import (
	"fmt"
	"nerversitup/services"
)

func main() {
	fmt.Println("Permutation input[ABCDA] result :", services.Permutation("ABC"))
	odd, times := services.FindOddInt([]int{7, 7, 7, 2, 2, 1})
	fmt.Println("FindOdd input[ 7, 7, 2, 2, 1] odd number is :", odd, " it appear :", times, " times")
	smileTimes := services.CountSmileys([]string{":]", ":[", ";*", ";~D"})
	fmt.Println("FindSlime input[ ", `":]", ":[", ";*", ";~D"`, " ] ", smileTimes, " times")
}
