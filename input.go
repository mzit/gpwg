package main

import "fmt"

func main() {
	str := "true"
	var ret bool
	switch str {
	case "true", "yes", "1":
		ret = true
	case "false", "no", "0":
		ret = false
	default:
		fmt.Println(str, "is not valid")
	}
	fmt.Println(ret)
}
