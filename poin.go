package main

import (
	"fmt"
)

func d(g *int) {
	fmt.Println(*g)
}

func main() {
	s := 5
	d(&s)
}
