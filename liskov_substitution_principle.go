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

func main() {
	rc := Rectangle{2, 3}
	UseIt(&rc)

}
