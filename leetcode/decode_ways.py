class Solution:
    def numDecodings(self, s: str) -> int:
        @cache
        def solve(s):
            if len(s) == 0: return 1
            if s[0] == '0': return 0
            if len(s) == 1: return 1

            if s[:2] > '26': 
                return solve(s[1:])            
            return solve(s[1:]) + solve(s[2:])
        
        return solve(s)
              
