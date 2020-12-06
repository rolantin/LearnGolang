package main

import "fmt"

type position struct {
	x, y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
	///
}

func WhereIsBadGuy(b badGuy) {
	x := b.name
	y := b.health
	z := b.pos.x
	w := b.pos.y
	fmt.Println(x, y, z, w)
}

func main() {
	var p position
	p.x = 2
	k := position{3, 4}
	fmt.Println(k)
	bad := badGuy{"rolan", 199, k}
	WhereIsBadGuy(bad)
}
