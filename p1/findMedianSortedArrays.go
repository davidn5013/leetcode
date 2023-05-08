package p1

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenN1, lenN2 := len(nums1), len(nums2)
	i, j := 0, 0
	res := make([]int, 0, lenN1+lenN2)

	for i < lenN1 && j < lenN2 {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	for i < lenN1 {
		res = append(res, nums1[i])
		i++
	}

	for j < lenN2 {
		res = append(res, nums2[j])
		j++
	}

	med := int((lenN1 + lenN2) / 2)
	if (lenN1+lenN2)%2 == 0 {
		return float64((res[med-1] + res[med])) / 2
	}
	return float64(res[med])
}
