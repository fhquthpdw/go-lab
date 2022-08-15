package main

import (
	rdsx "aws-demo/rds"
	"fmt"
)

func main() {
	if err := rdsx.ModifyPG(); err != nil {
		fmt.Println(err)
	}
}
