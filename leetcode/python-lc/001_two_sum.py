"""
001. Two Sum

This problem wasn't hard to do, but I realized that my solution is
not very optimal in a sense. I'm doing a log(n) algorithm, but I could
just do one pass.

For that, I'd do this:
1. Iterate over the list.
2. For each element, check if the difference between the target and
    the element is in the list. If it is, return the indexes.
3. If it isn't, add the element to a dictionary with the key being the
    element and the value being the index.
"""

from typing import List
from base import TestCase


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j]
        return []


if __name__ == "__main__":
    solution = Solution()

    testcase = TestCase("001", "twoSum", solution, [([2, 7, 11, 15], 9)], [[0, 1]])
    testcase.timeit()
