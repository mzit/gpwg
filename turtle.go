package main

import "fmt"

type locationTurtle struct {
	x, y int
}

func (l *locationTurtle) left() {
	l.x--
}

func (l *locationTurtle) right() {
	l.x++
}

func (l *locationTurtle) up() {
	l.y--
}

func (l *locationTurtle) down() {
	l.y++
}

func main() {
	var t locationTurtle
	t.up()
	t.up()
	t.left()
	t.left()
	fmt.Println(t)
	t.down()
	t.down()
	t.right()
	t.right()
	fmt.Println(t)
}
