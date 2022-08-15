package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// https://itnext.io/how-to-use-golang-generics-with-structs-8cabc9353d75
func main() {
	r := getCsvContent[Post](genPost())
	fmt.Println("+-+-+-+-+-+- Post Start +-+-+-+-+-+")
	fmt.Println(r[0].Id)
	fmt.Println(r[0].Title)
	fmt.Println(r[1].Id)
	fmt.Println(r[1].Title)
	fmt.Println("+-+-+-+-+-+- Post End +-+-+-+-+-+")

	r1 := getCsvContent[Category](genCategory())
	fmt.Println("+-+-+-+-+-+- Category Start +-+-+-+-+-+")
	fmt.Println(r1[0].Id)
	fmt.Println(r1[0].Name)
	fmt.Println(r1[1].Id)
	fmt.Println(r1[1].Name)
	fmt.Println("+-+-+-+-+-+- Category End +-+-+-+-+-+")

	// category
	category := Category{
		Id:   1,
		Name: "name",
	}
	cc := NewCache[Category]()
	cc.Set("category-key", category)

	// post
	post := Post{
		Id:    1,
		Title: "title",
	}
	cp := NewCache[Post]()
	cp.Set("post-key", post)
}

type Category struct {
	Id   int
	Name string
}

type Post struct {
	Id    int
	Title string
}

// dataAble limit cache data types
type dataType interface {
	Category | Post
}

type Cache[T dataType] struct {
	data map[string]T
}

func (c *Cache[T]) Set(k string, v T) {
	c.data[k] = v
}

func (c *Cache[T]) Get(k string) (v T) {
	if v, ok := c.data[k]; ok {
		return v
	}

	return
}

func NewCache[T dataType]() Cache[T] {
	c := Cache[T]{}
	c.data = make(map[string]T)

	return c
}

func getCsvContent[T dataType](dataSource []map[string]any) (r []T) {
	//pArr := []map[string]any{
	//	{"Id": 1, "Title": "title1"},
	//	{"Id": 2, "Title": "title2"},
	//}

	for _, item := range dataSource {
		p := new(T)
		_ = mapstructure.Decode(item, p)
		r = append(r, *p)
	}

	return
}

func genPost() []map[string]any {
	return []map[string]any{
		{"Id": 1, "Title": "title1"},
		{"Id": 2, "Title": "title2"},
	}
}

func genCategory() []map[string]any {
	return []map[string]any{
		{"Id": 91, "Name": "name1"},
		{"Id": 92, "Name": "name2"},
	}
}
