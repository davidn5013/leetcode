package grind75

import (
	"sort"
)

// Leetcode 15 3Sum Grind75 number 30 3Sum

// DistThreeSum return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
func DistThreeSum(nums []int) [][]int {
	res := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			low := i + 1
			high := len(nums) - 1
			sum := 0 - nums[i]

			for low < high {
				if nums[low]+nums[high] == sum {
					res = append(res, []int{nums[i], nums[low], nums[high]})
					for low < high && nums[low] == nums[low+1] {
						low++
					}
					for low < high && nums[high] == nums[high-1] {
						high--
					}
					low++
					high--
				} else if nums[low]+nums[high] > sum {
					high--
				} else {
					low++
				}
			}

		}
	}

	return res

}

/* My old slow soltion
// DistThreeSum return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
func DistThreeSum(nums []int) (res [][]int) {
	type t struct {
		i1 int
		i2 int
		i3 int
	}

	m := make(map[t]bool)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {

				if (i != j && i != k && j != k) &&
					nums[i]+nums[j]+nums[k] == 0 {
					temp := []int{nums[i], nums[j], nums[k]}
					sort.Ints(temp)

					if _, find := m[t{temp[0], temp[1], temp[2]}]; find == false {
						m[t{temp[0], temp[1], temp[2]}] = true
					}
				}

			}
		}
	}

	for v, _ := range m {
		res = append(res, []int{v.i1, v.i2, v.i3})
	}

	return res
}
*/

/*


func threeSum(nums []int) [][]int {
    res := [][]int{}

    sort.Ints(nums)

    for i := 0; i < len(nums)-2; i++ {
        if(i == 0 || (i > 0 && nums[i] != nums[i-1])) {
            low := i+1
            high := len(nums)-1
            sum := 0-nums[i]

            for (low < high) {
                if(nums[low] + nums[high] == sum) {
                    res = append(res, []int{nums[i], nums[low], nums[high]})
                    for (low < high && nums[low] == nums[low+1])
                        { low++ }
                    for(low < high && nums[high] == nums[high-1])
                        { high-- }
                    low++
                    high--
                } else if (nums[low] + nums[high] > sum) {
                    high--
                } else {
                    low++
                }
            }
        }
    }

    return res

}


func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := make([][]int, 0)
    // Bug?
    if len(nums) < 3 { return res }

    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 { break }
        if i > 0 && nums[i] == nums[i - 1] { continue }
        r := len(nums) - 1
        l := i + 1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum == 0 {
                res = append(res, []int{nums[i], nums[l], nums[r]})
                for l < r && nums[l] == nums[l+1] { l++ }
                for l < r && nums[r] == nums[r-1] { r-- }
                l++
                r--
            } else if sum > 0 {
                r--
            } else if sum < 0 {
                l++
            }
        }
    }
    return res
}

*/
