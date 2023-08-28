package codes

import "sort"

const (
	MaxNum = 1000
	MinNum = -1000
)

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	var sumMaps = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sumMaps[i] = target - nums[i]
	}

	var linkMaps = [MaxNum - MinNum + 1]int{}
	for i := 1; i < len(nums); i++ {
		v1, v2 := nums[i-1], nums[i]
		v := v1 + v2 / 2
		for j := v1; j <= v; j++ {
			linkMaps[j-MinNum] = i
		}

	}

}
