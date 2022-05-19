package main

import "fmt"

// Liskov Substitution Principle
// if you have some API that takes a base class and works properly,
// it should also work smoothly with the derived class

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}
func (r *Rectangle) GetHeight() int {
	return r.height
}
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected area of: ", expectedArea, ", but got: ", actualArea, "\n")
}

// Let's now create a Square that would be of type Sized too
type Square struct {
	Rectangle // it will force the idea of widht to be equal to height
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// The part that VIOLATES the Liscov Substitution Principle is that there are methods
// for setWidth setHeight that set both width and height
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}
func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

// The problem is that UseIt function is expecting something up the hierarchy (interface)
// as an argument. It should continue to work even if we proceed to extend objects
// (here we extended the rectangle and created a square object)

func main() {
	rc := Rectangle{2, 3}
	UseIt(&rc)

	sq := NewSquare(5)
	UseIt(sq)

}

// Expected area of: 20, but got: 20
// Expected area of: 50, but got: 100   --> because the call to setHeight not only set the height but also set the width
