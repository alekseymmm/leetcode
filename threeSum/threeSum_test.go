package threesum

import (
	"sort"
	"testing"
)

func calcExpectedThreeSum(nums []int) map[[3]int]bool {

	uniq := map[[3]int]bool{}

	ar := [3]int{}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					ar[0] = nums[i]
					ar[1] = nums[j]
					ar[2] = nums[k]
					sort.Ints(ar[:])
					uniq[ar] = true
				}
			}
		}
	}
	// res := make([][]int, len(uniq))
	// i := 0
	// for k := range uniq {
	// 	res[i] = []int{k[0], k[1], k[2]}
	// 	i++
	// }
	return uniq
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	expected := calcExpectedThreeSum(nums)
	result := ThreeSum(nums)

	if len(result) != len(expected) {
		t.Error("len of result", len(result),
			"does not match expected len", len(expected))
	}

	ar := [3]int{}
	for _, sl := range result {
		sort.Ints(sl)
		copy(ar[:], sl)
		if _, ok := expected[ar]; !ok {
			t.Error("wrong answer:", sl)
		}
	}
}
