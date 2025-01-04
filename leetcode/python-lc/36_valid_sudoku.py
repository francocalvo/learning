"""
36. Valid Sudoku

This was easy. Hardest part was determining the sqr index.
Also, I could've implemented this using sets, but the performance
seemed to be the same.
"""

from typing import Dict, List
from base import TestCase


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        row_hm: Dict[int, Dict[str, bool]] = {}
        col_hm: Dict[int, Dict[str, bool]] = {}
        sqr_hm: Dict[int, Dict[str, bool]] = {}

        for r in range(len(board)):
            row_hm[r] = {}
            for c in range(len(board[r])):
                if c not in col_hm:
                    col_hm[c] = {}

                s = board[r][c]
                if s == ".":
                    continue

                sqr = (r // 3) * 3 + c // 3

                if sqr not in sqr_hm:
                    sqr_hm[sqr] = {}

                if s in row_hm[r] or s in col_hm[c] or s in sqr_hm[sqr]:
                    return False

                row_hm[r][s] = True
                col_hm[c][s] = True
                sqr_hm[sqr][s] = True

        return True


if __name__ == "__main__":
    sk_1 = [
        ["5", "3", ".", ".", "7", ".", ".", ".", "."],
        ["6", ".", ".", "1", "9", "5", ".", ".", "."],
        [".", "9", "8", ".", ".", ".", ".", "6", "."],
        ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
        ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
        ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
        [".", "6", ".", ".", ".", ".", "2", "8", "."],
        [".", ".", ".", "4", "1", "9", ".", ".", "5"],
        [".", ".", ".", ".", "8", ".", ".", "7", "9"],
    ]

    sk_2 = [
        [".", ".", ".", ".", "5", ".", ".", "1", "."],
        [".", "4", ".", "3", ".", ".", ".", ".", "."],
        [".", ".", ".", ".", ".", "3", ".", ".", "1"],
        ["8", ".", ".", ".", ".", ".", ".", "2", "."],
        [".", ".", "2", ".", "7", ".", ".", ".", "."],
        [".", "1", "5", ".", ".", ".", ".", ".", "."],
        [".", ".", ".", ".", ".", "2", ".", ".", "."],
        [".", "2", ".", "9", ".", ".", ".", ".", "."],
        [".", ".", "4", ".", ".", ".", ".", ".", "."],
    ]

    sol = Solution()

    test = TestCase("36", "isValidSudoku", sol, [[sk_1, sk_2]], [True, False])
    test.test()
