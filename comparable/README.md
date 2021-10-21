## 如何比较一个结构体是否是空的
```go
package main

type User struct {
	Name string
}

func main() {
	u := User{}
	if u == (User{}) {  // 注意这里需要用一对()把 User{} 括起来
		//
	}
}
```