"""
238. Product of Array Except Self

This one was a bit tough. I knew I had to use prefix and suffix arrays, but
I was blanking on how they related to each other... Dumb, right?

Anyways, I first solved it with two arrays as that was easier to understand at
first, then modified it to use only one array, which complies with the O(1)
constraint.
"""

from typing import List
from base import TestCase


class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        ll = len(nums)
        res = [0] * ll

        px = 1
        for i in range(ll - 1, -1, -1):
            res[i] = nums[i] * px
            px = res[i]

        px = 1
        for i in range(1, ll):
            res[i - 1] = px * res[i]
            px = px * nums[i - 1]
        res[ll - 1] = px
        return res

    # First solution... I didn't do the O(1) constraint.
    # def productExceptSelf(self, nums: List[int]) -> List[int]:
    #     pf = [0 for n in range(len(nums))]
    #     sf = [0 for n in range(len(nums))]
    #
    #     ll = len(nums) - 1
    #     pf_x = 1
    #     sf_x = 1
    #     for i in range(len(nums)):
    #         pf[i] = pf_x * nums[i]
    #         sf[ll - i] = sf_x * nums[ll - i]
    #         pf_x = pf[i]
    #         sf_x = sf[ll - i]
    #
    #     for i in range(len(nums)):
    #         pf_x = pf[i - 1] if i - 1 >= 0 else 1
    #         sf_x = sf[i + 1] if i + 1 <= ll else 1
    #         print(f"pf_x: {pf_x}, sf_x: {sf_x}, i: {i}")
    #         nums[i] = pf_x * sf_x
    #
    #     return nums


if __name__ == "__main__":
    s = Solution()
    test = TestCase(
        "238",
        "productExceptSelf",
        s,
        [[[1, 2, 3, 4]], [[-1, 1, 0, -3, 3]], [[0, 0]]],
        [[24, 12, 8, 6], [0, 0, 9, 0, 0], [0, 0]],
    )

    test.test()
