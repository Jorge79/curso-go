package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i <= 15000; i++ {
		f, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))

		if err != nil {
			panic(err)
		}
		f.WriteString("Hello, World")
		i++
	}
}
