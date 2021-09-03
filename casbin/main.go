package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	modelFile := "model.conf"
	policyFile := "policy.conf"
	e, _ := casbin.NewEnforcer(modelFile, policyFile)

	sub := "article" // 用户
	obj := "data1"   // 资源
	act := "read"    // 操作
	//ok, err := e.Enforce(sub, obj, act)
	e.AddPolicy(sub, obj, act)
	p := e.GetPolicy()
	fmt.Println(p)
}
