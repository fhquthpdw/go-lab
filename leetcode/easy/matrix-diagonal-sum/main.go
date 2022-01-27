package main

import "fmt"

// 解题思路：遍历一维数组，每次遍历到的数字为对角线上的数字，求和即可
// - 需要考数组的长度
//   a. 如果是奇数，最终结果需要再减掉中间的数
//   b. 如果是偶数，最终结果不需要再减掉中间的数

func main() {
	sum1 := diagonalSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	sum2 := diagonalSum2([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println(sum1, sum2)
}

// 解法1:
func diagonalSum(mat [][]int) int {
	l := len(mat)

	sum := 0
	for i := 0; i < l; i++ {
		sum += mat[i][i] + mat[i][l-i-1]
	}

	if l%2 == 1 {
		sum -= mat[(l+1)/2-1][(l+1)/2-1]
	}

	return sum
}

// 解法2: 不新申明变量，直接用数组的第一个元素作为反回结果
func diagonalSum2(mat [][]int) int {
	l := len(mat)

	if l == 1 {
		return mat[0][0]
	}

	mat[0][0] = mat[0][0] + mat[0][l-1]
	for i := 1; i < l; i++ {
		mat[0][0] += mat[i][i] + mat[i][l-i-1]
	}

	if l%2 == 1 {
		mat[0][0] -= mat[(l+1)/2-1][(l+1)/2-1]
	}

	return mat[0][0]
}
