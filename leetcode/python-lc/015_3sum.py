"""
15. 3Sum
"""

from typing import Dict, List, Set, Tuple
from base import TestCase


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        visited: Set[int] = set()
        visited_pairs: Set[Tuple[int, int]] = set()
        sols: List[List[int]] = []

        for i, target in enumerate(nums):
            if target not in visited:
                visited.add(target)
                print("#### TARGET: ", target * -1)

                diffs: Dict[int, int] = {}
                for j, num in enumerate(nums):
                    print("Checking missing with ", num)
                    if (
                        j == i
                        or (target, num) in visited_pairs
                        or (num, target) in visited_pairs
                    ):
                        continue
                    else:
                        visited_pairs.add((target, num))
                        visited_pairs.add((num, target))

                    if num in diffs and diffs[num] != -1:
                        print("%%% Found: ", num, nums[diffs[num]])
                        sols.append([target, num, nums[diffs[num]]])
                        diffs[num] = -1
                    else:
                        print(f"Adding {-1*target-num} to diffs")
                        diffs[-1 * target - num] = j

        return sols


if __name__ == "__main__":
    sol = Solution()

    test = TestCase(
        "15",
        "threeSum",
        sol,
        [
            [[-1, 0, 1, 2, -1, -4]],
            # [[0, 1, 1]],
            # [[0, 0, 0]],
        ],
        [
            [[-1, -1, 2], [-1, 0, 1]],
            # [],
            # [[0, 0, 0]],
        ],
    )

    test.test()
