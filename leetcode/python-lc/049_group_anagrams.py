from typing import Dict, List
from base import TestCase


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        hm: Dict[str, List[str]] = {}

        for s in strs:
            s_sorted = "".join(sorted(s))
            if s_sorted not in hm:
                hm[s_sorted] = []
            hm[s_sorted].append(s)
        return list(hm.values())


if __name__ == "__main__":
    solution = Solution()

    testcase = TestCase(
        "049",
        "groupAnagrams",
        solution,
        [(["eat", "tea", "tan", "ate", "nat", "bat"],)],
        [[["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]],
    )
    testcase.timeit()
