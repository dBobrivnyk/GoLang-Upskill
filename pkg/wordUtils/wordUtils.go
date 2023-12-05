package wordUtils

func AreAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	sourceMap := make(map[rune]int)
	targetMap := make(map[rune]int)

	for _, letter := range a {
		sourceMap[letter]++
	}

	for _, letter := range b {
		targetMap[letter]++
	}

	// use maps.Equal
	for letter, sourceCount := range sourceMap {

		if targetCount, ok := targetMap[letter]; !ok || sourceCount != targetCount {
			return false
		}
	}

	return true
}

func reverse(s string) string {
	rev := []rune(s)

	for x, y := 0, len(rev)-1; x < len(rev)/2; x, y = x+1, y-1 {
		rev[x], rev[y] = rev[y], rev[x]
	}

	return string(rev)
}

func IsPalindrome(s string) bool {
	if s == reverse(s) {
		return true
	}
	return false
}
