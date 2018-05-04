package util

// GetLongestLineLength returns the length of the longest "line" within the
// argument string. For ex.:
//  GetLongestLineLength("Ghost!\nCome back here!\nRight now!") == 15
func GetLongestLineLength(s string) int {
	maxLength, currLength := 0, 0
	for _, c := range s {
		if c == '\n' {
			if currLength > maxLength {
				maxLength = currLength
			}
			currLength = 0
		} else {
			currLength++
		}
	}
	if currLength > maxLength {
		return currLength
	}
	return maxLength
}
