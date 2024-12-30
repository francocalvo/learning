"""
271. Encode and decode strings

This works OK. I guess another alternatives were: using the proposed thing on
NeetCode, or use something like Redis encoding.
"""

from base import TestCase
from typing import List


class Solution:
    def encode(self, strs: List[str]) -> str:
        if not strs:
            return ""
        return "#%m" + "#%m".join(strs)

    def decode(self, s: str) -> List[str]:
        return s.split("#%m")[1:]

    def test(self, strs: List[str]) -> List[str]:
        return self.decode(self.encode(strs))


if __name__ == "__main__":
    strs = [
        [["abc", "def", "ghi"]],
        [["abc", "def", "ghi", ""]],
        [[""]],
        [[]],
    ]

    s = Solution()
    test = TestCase("271", "test", s, strs, [s[0] for s in strs])

    test.test()
