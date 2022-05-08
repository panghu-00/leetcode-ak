package leetcode

func twoSum(nums []int, target int) []int {
	var numMap map[int]int = make(map[int]int)
	for index, value := range nums {
		numMap[value] = index
	}
	for index, value := range nums {
		if value, ok := numMap[target-value]; ok && value != index {
			return []int{index, value}
		}
	}
	return []int{}
}
