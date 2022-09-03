package greedyAlgorithm

func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0
	twenty := 0

	for _, bill := range bills {
		if bill == 5 {
			five++
			continue
		}

		if bill == 10 {
			if  five <= 0 {
				return false
			}
			five--
			ten++
			continue
		}

		if bill == 20 {
			if ten > 0 && five > 0 {
				ten--
				five--
				twenty++
				continue
			} else if five >= 3 {
				five -= 3
				twenty++
				continue
			} else {
				return false
			}
		}
	}
	return true
}
