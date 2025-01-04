"""125. Valid Palindrome

I first checked for a quick way to strip out all non-alphanumeric characters.
I found the re module and used the re.compile method to create a pattern that
matches all non-alphanumeric characters.

I think my second way is way more readable and easier to understand.

"""

from base import TestCase
import re


class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = "".join([c for c in s if c.isalnum()]).lower()
        return s == s[::-1]

    def isPalindrome_2(self, s: str) -> bool:
        pattern = re.compile(r"[\W_]+", re.ASCII)
        s = pattern.sub("", s).lower()

        start = 0
        end = len(s) - 1

        while start < end:
            if s[start] != s[end]:
                return False
            start += 1
            end -= 1

        return True


if __name__ == "__main__":
    sol = Solution()
    test = TestCase(
        "125",
        "isPalindrome",
        sol,
        [["A man, a plan, a canal: Panama"], [" "]],
        [True, True],
    )

    test.test()
