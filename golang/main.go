package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a t1 = 1
	var b t2 = 64
	fmt.Println(Join([]Stringer{a, b}))

	var c myFloat = 1.11

	P(c)
	P(1)
	P(1.2)
}

type myFloat float64

type Ordered interface {
	int | int64 | ~float64
}

func P[T Ordered](v T) {
	fmt.Println(v)
}

type Stringer interface {
	String() string
}

func Join[T Stringer](s []T) string {
	r := ""
	for _, v := range s {
		r = fmt.Sprintf("%s, %s", r, v.String())
	}
	return r
}

type t1 int
type t2 int64

func (i t1) String() string {
	return strconv.Itoa(int(i))
}

func (i t2) String() string {
	return strconv.Itoa(int(i))
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
