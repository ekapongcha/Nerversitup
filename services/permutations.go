package services

func Permutation(input string) []string {
	//In case of one string then return result
	if len(input) == 1 {
		return []string{input}
	}
	result := []string{}
	for i, str := range input {
		//get all input without str **note selected rune
		remaining := input[:i] + input[i+1:]
		//do all remain
		subPermutations := Permutation(remaining)

		//after get last remain then append it to str
		for _, subPermutations := range subPermutations {
			result = append(result, string(str)+subPermutations)
		}
	}
	return result
}
