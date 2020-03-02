package threesum

import "sort"

func twoSumSkip(nums []int, target int, skip int) [][]int {
	res := [][]int{}
	sums := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if i == skip {
			continue
		}
		cur := nums[i]
		if val, ok := sums[cur]; ok {
			res = append(res, []int{val, cur})
		}
		sums[target-cur] = cur
	}

	return res
}

func threeSum(nums []int) [][]int {
	uniq := map[[3]int]bool{}
	ar := [3]int{}
	for i, v := range nums {
		sums := twoSumSkip(nums, -v, i)
		for _, pair := range sums {
			ar[0] = v
			ar[1] = pair[0]
			ar[2] = pair[1]
			sort.Ints(ar[:])
			uniq[ar] = true
		}
	}
	res := [][]int{}
	for triple := range uniq {
		sl := []int{triple[0], triple[1], triple[2]}
		res = append(res, sl)
	}
	return res
}

func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}
