package main

// https://www.ardanlabs.com/blog/2021/04/using-bitmasks-in-go.html
import (
	"fmt"
	"strings"
)

type KeySet byte

const (
	Copper KeySet = 1 << iota
	Jade
	Crystal
	maxKey
)

func (k KeySet) String() string {
	if k >= maxKey {
		return fmt.Sprintf("<unknown key: %d>", k)
	}

	switch k {
	case Copper:
		return "copper"
	case Jade:
		return "jade"
	case Crystal:
		return "crystal"
	}

	var names []string
	for key := Copper; key < maxKey; key <<= 1 {
		if k&key != 0 {
			names = append(names, key.String())
		}
	}

	return strings.Join(names, "|")
}

type Player struct {
	Name string
	Keys KeySet
}

func (p *Player) AddKey(key KeySet) {
	p.Keys |= key
}

func (p *Player) HasKey(key KeySet) bool {
	return p.Keys&key != 0
}

func (p *Player) RemoveKey(key KeySet) {
	p.Keys &= ^key
}
