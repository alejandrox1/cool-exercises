package main

// twoSum works if the array is sorted.
func twoSum(values []int, target int) (int, int) {
	if len(values) <= 1 {
		return -1, -1
	}

	for i, j := 0, len(values)-1; i < len(values); {
		if i == j || i > j {
			return -1, -1
		}

		sum := values[i] + values[j]

		if sum == target {
			return i, j
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

	return -1, -1
}

// twoSumUnsorted goes through all elements checking if the difference between
// the target and the current value has already been seen (if current and a
// previous value sum equal to the target).
func twoSumUnsorted(values []int, target int) (int, int) {
	if len(values) <= 1 {
		return -1, -1
	}

	seen := make(map[int]int)
	for i, value := range values {
		if j, ok := seen[value]; ok {
			return i, j
		}

		seen[target-value] = i
	}

	return -1, -1
}
