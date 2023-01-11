package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// https://itnext.io/how-to-use-golang-generics-with-structs-8cabc9353d75
func main() {
	r := getCsvContent[Post](genPost())
	fmt.Printf("%#v\n", r)

	r1 := getCsvContent[Category](genCategory())
	fmt.Printf("%#v", r1)

	// category
	category := Category{
		ID:   1,
		Name: "name",
	}
	cc := NewCache[Category]()
	cc.Set("category-key", category)

	// post
	post := Post{
		ID:    1,
		Title: "title",
	}
	cp := NewCache[Post]()
	cp.Set("post-key", post)
}

type Category struct {
	ID   int64
	Name string
}

func (c Category) GetId() int64 {
	return c.ID
}

type Post struct {
	ID    int64
	Title string
}

func (c Post) GetId() int64 {
	return c.ID
}

// dataAble limit cache data types
type dataType interface {
	Category | Post
	GetId() int64
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

// IdMap 泛型应用
func IdMap[T dataType](source []T) (r map[int64]T) {
	for _, item := range source {
		r[item.GetId()] = item
	}

	return
}
