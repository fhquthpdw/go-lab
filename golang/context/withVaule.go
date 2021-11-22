package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "a", "aa")
	ctx2 := context.WithValue(ctx, "b", "bb")
	Get(ctx2, "a")
	Get(ctx2, "b")
	Get(ctx2, "c")
}

func Get(ctx context.Context, k interface{}) {
	if v, ok := ctx.Value(k).(string); ok {
		fmt.Println(v)
	}

}
