package main

import "fmt"

func main() {
	twoWorlds := terraform1("New", "aaa", "bbbb", "cccc")
	fmt.Println(twoWorlds)

	planets := []string{
		"sajhdjka", "hjkajkhs", "hjasdhj", "jdjakd",
	}
	newPlanets := terraform1("new", planets...)
	fmt.Println(newPlanets)
}

func terraform1(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	for i := range newWorlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}
