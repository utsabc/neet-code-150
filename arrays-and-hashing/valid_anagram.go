package arraysandhashing

import (
	"slices"
	"strings"
)

func ValidAnagramSlow(s string, t string) bool {

	sCharArr := strings.Split(s, "")
	tCharArr := strings.Split(t, "")

	if len(sCharArr) != len(tCharArr) {
		return false
	}

	slices.Sort(sCharArr)
	slices.Sort(tCharArr)

	for i, _ := range sCharArr {
		if sCharArr[i] != tCharArr[i] {
			return false
		}
	}
	return true
}

func ValidAnagramFast(s string, t string) bool {

	sCharArr := strings.Split(s, "")
	tCharArr := strings.Split(t, "")

	if len(sCharArr) != len(tCharArr) {
		return false
	}

	stateMapS := make(map[string]int)
	stateMapT := make(map[string]int)

	for i := 0; i < len(sCharArr); i++ {
		stateMapS[sCharArr[i]] = stateMapS[sCharArr[i]] + 1
		stateMapT[tCharArr[i]] = stateMapT[tCharArr[i]] + 1
	}

	for k, v := range stateMapS {
		if v != stateMapT[k] {
			return false
		}
	}

	return true
}
