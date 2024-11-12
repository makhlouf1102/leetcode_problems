package easy

import (
	"fmt"
	"strings"
)

func countBits(n int) []int {
	ans := make([]int, n+1)
	var bin string
	for i := 0; i <= n; i++ {
		bin = fmt.Sprintf("%b", i)
		ans[i] = strings.Count(bin, "1")
	}

	return ans
}
