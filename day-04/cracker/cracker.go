package cracker

import (
	"strconv"
	"strings"
)

// CountPasswords Following the rules laid out, count the passwords in the range defined by low andhigh boundary
func CountPasswords(lowBoundary int, highBoundary int) (int, error) {
	count := 0
	for i := lowBoundary; i <= highBoundary; i++ {
		digits, err := splitDigits(i)
		if err != nil {
			return 0, err
		}

		if checkLength(digits, 6) && checkRange(i, lowBoundary, highBoundary) && checkDoubleDigits(digits) && checkNoDescending(digits) {
			count++
		}
	}

	return count, nil
}

func splitDigits(toSplit int) ([]int, error) {
	var retVal []int
	stringArray := strings.Split(strconv.Itoa(toSplit), "")

	for _, digitChar := range stringArray {
		temp, err := strconv.Atoi(digitChar)

		if err != nil {
			return nil, err
		}
		retVal = append(retVal, temp)
	}

	return retVal, nil
}

func joinDigits(toJoin []int) (int, error) {
	var stringDigits []string

	for _, digit := range toJoin {
		stringDigits = append(stringDigits, strconv.Itoa(digit))
	}

	stringNumber := strings.Join(stringDigits, "")

	return strconv.Atoi(stringNumber)
}

func checkLength(candidate []int, requiredLength int) bool {
	return len(candidate) == requiredLength
}

func checkRange(candidate int, lowBoundary int, highBoundary int) bool {
	return candidate >= lowBoundary && candidate <= highBoundary
}

func checkDoubleDigits(candidate []int) bool {
	for i := 1; i < len(candidate); i++ {
		if candidate[i] == candidate[i-1] {
			return true
		}
	}

	return false
}

func checkNoDescending(candidate []int) bool {
	for i := 1; i < len(candidate); i++ {
		if candidate[i] < candidate[i-1] {
			return false
		}
	}

	return true
}
