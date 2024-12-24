"""
217: Contains duplicates
"""

from typing import Dict, List
from base import TestCase


class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        hm: Dict[int, bool] = {}

        for n in nums:
            if hm.get(n, False):
                return True
            hm[n] = True

        return False


if __name__ == "__main__":
    solution = Solution()

    test = TestCase(
        "217",
        "containsDuplicate",
        solution,
        [
            [[1, 2, 3, 1]],
            [[1, 2, 3, 4]],
            [[1, 1, 1, 3, 3, 4, 3, 2, 4, 2]],
        ],
        [True, False, True],
    )

    test.timeit()
