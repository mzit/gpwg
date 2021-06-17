package main

import (
	"fmt"
	"strings"
)

func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	planets := []string{" aaa ", " bbb ", " ccc "}
	hyperspace(planets)
	fmt.Println(planets)
	fmt.Println(strings.Join(planets, ""))
}
