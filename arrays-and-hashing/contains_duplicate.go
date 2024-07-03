package arraysandhashing

import "slices"

func FindDuplicatesSlow(nums []int) bool {

	slices.Sort(nums) // nLogn

	for index, currentElement := range nums {
		var nextElement int
		if index < len(nums)-1 {
			nextElement = nums[index+1]
			if currentElement == nextElement {
				return true
			}
		}
	}
	return false
}

func FindDuplciatesFast(nums []int) bool {

	stateMap := make(map[int]bool)

	for _, value := range nums {
		if stateMap[value] == true {
			return true
		}
		stateMap[value] = true
	}
	return false

}
