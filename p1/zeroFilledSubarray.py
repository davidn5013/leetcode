Class Solution:
    def zeroFilledSubarray(self, nums: List[int]) -> int:
        res,zcnt = 0,0
        for i in range(len(nums)):
            if nums[i] == 0:
                zcnt += 1
            else:
                zcnt = 0
            res += zcnt
        return res
