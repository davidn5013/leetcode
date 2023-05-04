package lc26

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		ret[v] = struct{}{}
	}
	return len(ret)
}
