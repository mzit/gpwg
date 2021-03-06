package main

import "fmt"

type stats struct {
	level             int
	endurance, health int
}

func levelup(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

func main() {

	player := character{name: "aaaaa"}
	levelup(&player.stats)
	fmt.Printf("%+v\n", player)
}
