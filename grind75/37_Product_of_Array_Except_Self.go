package grind75

// Leetcode 238. Product of Array Except Self
// Grind75 Number 37

/*
Given an integer array nums, return an array answer such that answer[i] is
equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/

/* This is a O(N) solution */

// ProductExceptSelf return an array answer such that answer[i]
// is equal to the product of all the elements of nums except nums[i].
func ProductExceptSelf(nums []int) (res []int) {
	n := len(nums)
	// TODO better names for lhs, rhs
	lhs, rhs := 1, 1
	res = make([]int, n)

	// Prepping res with 1, remember to change this if you add och subtracting
	for i := 0; i < n; i++ {
		res[i] = 1
	}

	// This Variation while do left and right
	// instead first filling res with value instead doing this
	// res[0]=1;for i:=1;i<n;i++ {res[i]=res[i-1]*nums[i-1];}
	// and then doing
	// prod:=1;for i:=n-1;i>= 0;i-- {res[i]=res[i]*prod;prod*=nums[i];}
	// this is filling res from the left while doing the product from the right
	// both version are O(N) but this save some calculation time
	for i := 0; i < n; i++ {
		// TODO I don't get this algorithm fully so need get back to it
		// log.Println(res)
		res[i] *= lhs
		res[n-i-1] *= rhs
		lhs *= nums[i]
		rhs *= nums[n-i-1]
		// log.Println(res)
	}

	return res
}
