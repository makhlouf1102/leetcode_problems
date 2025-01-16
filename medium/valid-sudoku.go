package medium

func isValidSudoku(board [][]byte) bool {

	var cols, rows, squares [9][9]bool
	for i, r := range board {
		for j, v := range r {
			if v != '.' {
				k := int(board[i][j]-'0') - 1
				if rows[i][k] || cols[j][k] || squares[i/3*3+j/3][k] {
					return false
				}

				rows[i][k], cols[j][k], squares[i/3*3+j/3][k] = true, true, true

			}
		}

	}
	return true

}
