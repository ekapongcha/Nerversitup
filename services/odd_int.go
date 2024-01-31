package services

// input []int
// outpunt int (odd int) , int (times)
func FindOddInt(input []int) (int, int) {
	//declare map
	countingMap := map[int]int{}
	for _, item := range input {
		_, ok := countingMap[item]
		if ok {
			count := countingMap[item]
			count = count + 1
			countingMap[item] = count
		} else {
			countingMap[item] = 1
		}
	}

	index := 0
	var mostAppear int
	// loop find smallest times
	for key, val := range countingMap {
		if index == 0 {
			mostAppear = key
			index++
			continue
		} else {
			if countingMap[mostAppear] > val {
				mostAppear = key
				continue
			}
		}
	}

	return mostAppear, countingMap[mostAppear]
}
