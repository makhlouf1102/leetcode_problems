package easy

func twoSum(nums []int, target int) []int {
	subs := make(map[int]int)

	for i, num := range nums {
		if j, ok := subs[num]; ok {
			return []int{j, i}
		}
		subs[target-num] = i
	}

	return nil
}
