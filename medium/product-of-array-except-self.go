package medium

func productExceptSelf(nums []int) []int {
	var n int = len(nums)
	var res []int = make([]int, n)
	var prefix int = 1
	var postfix int = 1
	var j int = 0
	for i := 0; i < n; i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	for i := range n {
		j = n - i - 1
		res[j] *= postfix
		postfix *= nums[j]
	}
	return res

}
