package main

import "fmt"

// 解题思路: 遍历一个数组，在另一个数组中查找当前相同的元素，找到之后，写到入返回值的数组中，并删除这个元素

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1, 3}, []int{2, 2}))
}

func intersect(nums1 []int, nums2 []int) []int {
	r := make([]int, 0)
	for _, n1 := range nums1 {
		for j, n2 := range nums2 {
			if n1 == n2 {
				r = append(r, n2)
				copy(nums2[j:], nums2[j+1:])
				nums2 = nums2[:len(nums2)-1]
				break
			}
		}
	}

	return r
}
