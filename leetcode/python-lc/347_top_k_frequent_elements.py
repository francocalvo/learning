"""
347. Top K Frequent Elements

This works and ranks 92% in terms of speed, but I know there is a quick select
solution that can be used to solve this
"""

from typing import Dict, List
from base import TestCase


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        hm: Dict[int, int] = {}
        for num in nums:
            if num not in hm:
                hm[num] = 0

            hm[num] += 1

        zip_list = list(hm.items())
        zip_list.sort(key=lambda x: x[1], reverse=True)

        return [x[0] for x in zip_list[:k]]


if __name__ == "__main__":
    solution = Solution()

    inputs = [([1, 1, 1, 2, 2, 3], 2), ([1], 1), ([3, 0, 1, 0], 1)]
    expected = [[1, 2], [1], [0]]

    testcase = TestCase("347", "topKFrequent", solution, inputs, expected)

    testcase.test()
