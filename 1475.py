from typing import List
class Solution:
    def finalPrices(self, prices: List[int]) -> List[int]:
        n = len(prices)
        stack = [0] 
        ans = [0] * n

        for i in range(n):
            while stack and prices[stack[-1]] >= prices[i]:
                idx = stack.pop()
                ans[idx] = prices[idx] - prices[i] 
            stack.append(i)

        while stack:
            r = stack.pop()
            ans[r] = prices[r]
        
        return ans
def main():
    prices = [8,4,6,2,3]
    solution = Solution()
    print(solution.finalPrices(prices))

if __name__ == "__main__":
    main()
