package main

import "fmt"

func main() {
	fmt.Println(modifyString("???c"))
}

func modifyString(s string) string {
	r := make([]byte, 0, len(s))

	qNum := 0
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v == '?' {
			qNum++
			continue
		}

		for j := 0; j < qNum; j++ {
			idx := i - j - 1 // 当前问号的下标

			if j == 0 { // 是不是问号区间的第一个问号
				before := v
				if before == 'a' {
					if idx > 0 {
						r = append(r, before-1)
					} else {
						r = append(r, 'z')
					}
				}
			}

			/*
				iv := s[i-j]
				if iv == 'a' {
					r = append(r, 'z')
				} else {
					r = append(r, iv-1)
				}
			*/
		}

		r = append(r, v)
	}

	return string(r)
}
