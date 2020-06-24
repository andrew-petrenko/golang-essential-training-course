package structs

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}
