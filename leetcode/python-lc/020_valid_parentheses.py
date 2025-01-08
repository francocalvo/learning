"""
20. Valid parentheses

Easy easy stack problem. I used a hashmap as I didn't want many ifs...
"""

from typing import List
from base import TestCase


class Solution:
    def isValid(self, s: str) -> bool:
        stack: List[str] = []
        reverse = {")": "(", "]": "[", "}": "{"}

        for char in s:
            if char in ("(", "{", "["):
                stack.append(char)
            else:
                if len(stack) == 0 or reverse[char] != stack.pop():
                    return False

        return len(stack) == 0


if __name__ == "__main__":
    sol = Solution()
    test = TestCase(
        "20",
        "isValid",
        sol,
        [
            ["()"],
            ["()[]{}"],
            ["(]"],
            ["([)]"],
            ["{[]}"],
            ["]"],
        ],
        [True, True, False, False, True, False],
    )

    test.test()
