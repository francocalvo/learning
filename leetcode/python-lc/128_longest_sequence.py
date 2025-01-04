"""
128. Longest Consecutive Sequence

I tried to look each item and start creating and merging lists, but it wasn't
fast enough to pass the test cases. So, I looked at the hints on NeetCode and
created this.

The set construction is O(n) and the while loop is O(n) as well. So, the overall
complexity is O(n). This is because we are going to visit each element at most
twice.
"""

from typing import List
from base import TestCase


class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        hset = set(nums)
        m_l = 0

        for num in nums:
            if num - 1 not in hset:
                count = 1
                curr = num + 1
                while curr in hset:
                    count += 1
                    curr += 1

                if count > m_l:
                    m_l = count

        return m_l


if __name__ == "__main__":
    sol = Solution()

    test = TestCase(
        "128",
        "longestConsecutive",
        sol,
        [
            [[0, 3, 7, 2, 5, 8, 4, 6, 0, 1]],
            [[0, 0, -1]],
            [[9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6]],
        ],
        [9, 2, 7],
    )

    test.test()

    # ops: List[List[int]] = []
    # for n in nums:
    #     # print(f"n: {n}, ops: {ops}")
    #     merge = False
    #     merged = False
    #     base_idx = 0
    #     for i in range(len(ops)):
    #         # print(f"i: {i}, ops[i]: {ops[i]}, n: {n}")
    #         if n >= ops[i][0] and n <= ops[i][-1]:
    #             merged = True
    #             break
    #         # print(f"n: {n}, ops[i][-1]: {ops[i][-1]}, ops[i][0]: {ops[i][0]}")
    #         if n == ops[i][-1] + 1 or n == ops[i][0] - 1:
    #             # print("match found")
    #             if merge:
    #                 # print("merging")
    #                 if n == ops[i][-1] + 1:
    #                     ops[i].append(n)
    #                     ops[i] = ops[i] + ops[base_idx]
    #                     ops.pop(base_idx)
    #                 else:
    #                     ops[base_idx].append(n)
    #                     ops[i] = ops[base_idx] + ops[i]
    #                     ops.pop(base_idx)
    #                 merged = True
    #                 break
    #             merge = True
    #             base_idx = i
    #
    #     if merged:
    #         continue
    #
    #     if merge:
    #         if n == ops[base_idx][-1] + 1:
    #             ops[base_idx] = ops[base_idx] + [n]
    #         else:
    #             ops[base_idx] = [n] + ops[base_idx]
    #         continue
    #
    #     ops.append([n])
    #
    # # print("%%%%%%%")
    # # print()
    # # print()
    # # print(ops)
    # mx_l = 0
    # for o in ops:
    #     if len(o) > mx_l:
    #         mx_l = len(o)
    # return mx_l
