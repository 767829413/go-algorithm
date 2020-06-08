package backtrack

import "fmt"

//下标表示行,值表示queen存储在哪一列
var result [8]int

func Set8Queen(row int) {
	if row == 8 { //8个棋子都放置好了，打印结果
		printQueen()
		return
	}
	for column := 0; column < 8; column++ { //每一行对应8列结果
		if isOK(row, column) {
			result[row] = column
			Set8Queen(row + 1)
		}
	}
}

//判断row行column列放置是否合适
func isOK(row, column int) bool {
	leftUp, rightUp := column-1, column+1
	for i := row - 1; i >= 0; i-- { //逐行往上考察每一行
		//先判断第i行的column列是否放置
		if result[i] == column {
			return false
		}
		//判断左上对角线：第i行leftup列是否有放置
		if leftUp >= 0 && result[i] == leftUp {
			return false
		}
		//判断右上对角线：第i行rightup列是否有放置
		if rightUp < 8 && result[i] == rightUp {
			return false
		}
		leftUp--
		rightUp++
	}
	return true
}

func printQueen() {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if result[row] == column {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
