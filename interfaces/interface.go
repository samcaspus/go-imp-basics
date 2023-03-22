package interfaces

import "fmt"

type shape interface {
	edges() int
	colors() string
}

type square struct {
	edge  int
	color string
}

type circle struct {
	edge  int
	color string
}

func (s *square) edges() int {
	return s.edge
}

func (s *square) colors() string {
	return s.color
}

func (c *circle) edges() int {
	return c.edge
}

func (c *circle) colors() string {
	return c.color
}

func main() {
	s := square{4, "red"}
	c := circle{0, "blue"}
	shapes := []shape{&s, &c}
	for _, shape := range shapes {
		fmt.Println(shape.edges())
		fmt.Println(shape.colors())
	}
}
