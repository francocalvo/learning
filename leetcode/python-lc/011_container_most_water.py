"""
11. Container With Most Water

This was super easy knowing the two pointer approach.
I could add a short circuit to break the loop if the container height is the max
height in the list.
"""

from typing import List
from base import TestCase


class Solution:
    def maxArea(self, height: List[int]) -> int:
        water = 0
        left = 0
        right = len(height) - 1
        max_height = max(height)
        while left < right:
            container_height = min(height[left], height[right])
            water = max(water, (right - left) * container_height)
            if container_height == max_height:
                break
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        return water


if __name__ == "__main__":
    sol = Solution()
    test = TestCase(
        "11",
        "maxArea",
        sol,
        [
            [[1, 8, 6, 2, 5, 4, 8, 3, 7]],
            [[1, 1]],
            [[4, 3, 2, 1, 4]],
            [[1, 2, 1]],
        ],
        [49, 1, 16, 2],
    )

    test.test()
