package simple_factory

type Shape interface {
	Draw() string
}

func NewShape(style int) Shape {
	if style == 1 {
		return &Circle{name: "circle"}
	} else if style == 2 {
		return &Rectangle{name: "rectangle"}
	} else {
		return &Square{name: "square"}
	}
}

type Circle struct {
	name string
}

func (s Circle) Draw() string {
	return "circle"
}

type Rectangle struct {
	name string
}

func (s Rectangle) Draw() string {
	return "rectangle"
}

type Square struct {
	name string
}

func (s Square) Draw() string {
	return "square"
}
