import timeit


class TestCase:
    def __init__(self, tag, name, solution, input, expected):
        self.tag = tag
        self.name = name
        self.solution = solution
        self.input = input
        self.expected = expected

    def execute(self, method):
        for i, e in zip(self.input, self.expected):
            assert method(*i) == e

    def test(self):
        method = getattr(self.solution, self.name)
        for i, e in zip(self.input, self.expected):
            print(f"Input: {i}. Expected: {e}.")
            res = method(*i)
            if res == e:
                print("Passed.")
            else:
                print("Failed. Answer: ", res)

    def timeit(self):
        method = getattr(self.solution, self.name)
        t = timeit.timeit(lambda: self.execute(method), number=1000)
        print(f"{self.tag}. {self.name}: {t:.6f} seconds")
