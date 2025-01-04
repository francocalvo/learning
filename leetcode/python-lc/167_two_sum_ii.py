"""
167. Two Sum II - Input array is sorted

Ok, I first did the msot naive thing, going from the start for both pointers,
but this makes it longer than needed.
This method takes end and start, and dependending where we are relative to the
target, it moves one pointer. It beat 100% o_O
"""

from typing import List


class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        start = 0
        end = len(numbers) - 1

        while start < end:
            sum = numbers[start] + numbers[end]
            if sum > target:
                end -= 1
            elif sum < target:
                start += 1
            else:
                return [start + 1, end + 1]

        return [0, 0]
