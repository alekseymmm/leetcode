package threesum

func twoSumSkip(nums []int, target int, skip int) []int {
	res := []int{-1, -1}
	sums := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if i == skip {
			continue
		}
		cur := nums[i]
		if val, ok := sums[cur]; ok {
			res[0] = val
			res[1] = i
			return []int{val, i}
		}
		sums[target-cur] = i
	}

	return res
}

func threeSum(nums []int) [][]int {

	return [][]int{}
}

func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}
