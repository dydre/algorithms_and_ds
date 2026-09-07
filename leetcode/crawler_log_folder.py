from typing import List

class Solution:
    def minOperations(self, logs: List[str]) -> int:
        stack = []
        for log in logs:
            if log == "../":
                if stack:
                    stack.pop()
            elif log == "./":
                continue
            else:
                stack.append(log)
        return len(stack)

def main():
    solution = Solution()

    log1 = ["d1/", "d2/", "./", "d3/", "../", "d31/"]
    print(solution.minOperations(log1))

if __name__ == "__main__":
    main()
