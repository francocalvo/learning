"""
242: Valid Anagram
Neetcode: "Arrays & Hashing"
"""

from typing import Dict
from base import TestCase


class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        hm: Dict[str, int] = {}
        for char in s:
            hm[char] = hm.get(char, 0) + 1

        for char in t:
            val = hm.get(char, 0)
            if val == 0:
                return False
            hm[char] = val - 1

        return True


if __name__ == "__main__":
    TestCase(
        "242",
        "isAnagram",
        Solution(),
        [["anagram", "nagaram"], ["rat", "car"]],
        [True, False],
    ).timeit()
