## 如何给一个函数添加一个 function 类型的接收器
```go
package main

import "fmt"

type FType func() string

func (f FType) P() {
	fmt.Println(f())
}

func main() {
	WayOne()
	WayTwo()
}

func WayOne() {
	var f FType
	f = func() string {
		return "this is from way one"
	}
	f.P()
}

func WayTwo() {
	f := FType(func() string {
		return "this is from way two"
	})
	f.P()
}
```

## 函数的接收器可以是一个类型的指针但是不可以是一个指针类型
```go
// 函数的接收器是一个类型的指针 
type User struct {
    Name string
}

func (u *User) GetName() string {
    return u.Name
}
```

```go
// 函数的接收器是一个指针类型
/** 编译不通过 **/
type T *int

func (t T) GetName() bool {
	return true
}
```