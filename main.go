package main

import "fmt"

func main() {
	all := map[int64][]int64{
		1:  {2, 3, 4},
		5:  {6, 7, 8},
		4:  {9, 10, 11},
		7:  {12, 13, 14},
		11: {15, 16, 17},
		17: {18, 19},
		19: {20, 21},
		21: {22},
	}

	o := getChildRoles2([]int64{1, 5}, all)
	fmt.Println(o)
}

func getChildRoles2(mainIds []int64, mainChildArr map[int64][]int64) (childRoleIdArr []int64) {
	for _, mainId := range mainIds {
		childRoleIdArr = append(childRoleIdArr, mainId)
		roleIdArr, ok := mainChildArr[mainId]
		if !ok {
			continue
		}

		childRoleIdArr = append(childRoleIdArr, getChildRoles2(roleIdArr, mainChildArr)...)
	}
	return
}
