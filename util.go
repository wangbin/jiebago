package jiebago

import (
	"regexp"
)

// Split sentence using regular expression.
func RegexpSplit(r *regexp.Regexp, sentence string) chan string {
	result := make(chan string)
	go func() {
		locs := r.FindAllStringIndex(sentence, -1)
		lastLoc := 0
		for _, loc := range locs {
			if loc[0] == lastLoc {
				result <- sentence[loc[0]:loc[1]]
			} else {
				result <- sentence[lastLoc:loc[0]]
				result <- sentence[loc[0]:loc[1]]
			}
			lastLoc = loc[1]
		}
		if lastLoc < len(sentence) {
			result <- sentence[lastLoc:]
		}
		close(result)
	}()
	return result
}
