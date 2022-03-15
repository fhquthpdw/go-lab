package main

func main() {
	s := []int{1, 2, 3, 4, 5}
	move(s)
}

func move(arr []int) {
	for i, j := 0, len(arr)-1; i != j && i < j; {
		if arr[i]%2 == 0 { // 前面游标是偶数
			i++
			if arr[j]%2 == 0 { // 后面游标是偶数
			} else {
				j--
			}
		} else { // 前面游标是奇数
			if arr[j]%2 == 0 { // 后面游标是偶数	交换
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			} else { // 后面游标是奇数 不交换 移动后面的游标
				j--
			}
		}
	}
}
