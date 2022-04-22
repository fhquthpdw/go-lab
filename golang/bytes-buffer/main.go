package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	buf1 := bytes.NewBuffer([]byte{})
	buf1.Write([]byte("hello world!"))
	buf1.Write([]byte("\nFrom iHerb"))
	buf1.WriteString("\nfrom Daochun")
	buf1.WriteByte('b')

	// Read
	/* var b = make([]byte, 0, 2)
	for {
		if n, err := buf1.Read(b); err != io.EOF && n > 0 {
			fmt.Println(n)
			fmt.Printf("%s", string(b))
		} else {
			break
		}
	}*/

	// ReadByte
	for {
		if k, err := buf1.ReadByte(); err != io.EOF {
			fmt.Printf("%s", string(k))
		} else {
			fmt.Printf("\nEOF\n")
			break
		}
	}
}
