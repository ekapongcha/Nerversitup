package services

func CountSmileys(input []string) int {
	if len(input) == 0 {
		return 0
	}

	validPattern := []string{
		":D",
		";D",
		":-D",
		":~D",
		";-D",
		";~D",
		":)",
		";)",
		":-)",
		":~)",
		";-)",
		";~)",
	}
	count := 0
	for _, item := range input {
		if contains(validPattern, item) {
			count++
		}
	}

	return count
}

func contains(pattern []string, input string) bool {
	for _, a := range pattern {
		if a == input {
			return true
		}
	}
	return false
}
