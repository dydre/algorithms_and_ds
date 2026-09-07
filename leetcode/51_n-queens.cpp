class Solution {
public:
    vector<vector<string>> res;

    void backtrack(vector<string>& board, int row) {
        int n = board.size();
        if (row == n) {
            res.push_back(board);
            return ;
        }
        for (int col = 0; col < n; col++) {
            if (notHanging(board, row, col)) {
                board[row][col] = 'Q';
                backtrack(board, row + 1);
                board[row][col] = '.';
            }
        }
    }

    bool notHanging(vector<string>& board, int row, int col) {
        int n = board.size();
        for (int i = 0; i < row; i++) {
            if (board[i][col] == 'Q') {
                return false;
            }
        }
        for (int i = 1; i <= min(row, col); i++) {
            if (board[row - i][col - i] == 'Q') {
                return false;
            }
        }
        for (int i = 1; i <= min(row, n - 1 - col); i++) {
            if (board[row - i][col + i] == 'Q') {
                return false;
            }
        }
        return true;
    }
    vector<vector<string>> solveNQueens(int n) {
        vector<string> board(n, string(n, '.'));    
        backtrack(board, 0);
        return res;
    }
};
