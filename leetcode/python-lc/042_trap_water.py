"""
42. Trap water

Ok, I really flew through this one, less than 20min.
What this do is: 
1. For each pair of walls, find the minimum height. If the minimum height is 
   greater than the current height, then we trapped some water. 
2. Move the lower wall to the next wall.
3. Substract the space taken by the lower wall from the water trapped. If the
   lower wall height is greater than the current height, then we only substract
   up to the current height.
"""

from typing import List
from base import TestCase


class Solution:
    def trap(self, height: List[int]) -> int:
        curr_height = 0
        start = 0
        end = len(height) - 1
        water = 0

        while start < end:
            min_height = min(height[start], height[end])
            length = (end - start) - 1
            if min_height > curr_height:
                water += (min_height - curr_height) * length

            if min_height == height[start]:
                start += 1
            else:
                end -= 1
            water -= min(curr_height, min_height)
            curr_height = max(curr_height, min_height)

        return water


if __name__ == "__main__":
    sol = Solution()
    test = TestCase(
        "42",
        "trap",
        sol,
        [
            [[0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]],
            [[4, 2, 0, 3, 2, 5]],
        ],
        [6, 9],
    )

    test.test()
