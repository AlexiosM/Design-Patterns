package main

import "fmt"

// OCP
// open for extention, but closed for modification
// Specification
type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// func (f *Filter) FilterBySizeAndColor ....
// by adding more functionality for each probable case
// we lead to violation of the Open Close Principle, because if we go back
// and modify Product methods and add more things, we should also change all Filter methods,
// so we are interfering with something already been written and tested.
// Solution comes with: Specification interface

// The interface type is Open for extension
// (each new product attribute we want to filter we simply and a new type of this interface)
// However it is Closed for modification, you are unlikely to ever modify betterfilter. It
// is already very flexible. Already created structures and methods do not need to be modified.
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

// A Composite Specification is a combinator that combines two different specifications
type AndSpecification struct {
	first, second Specification
}

// So here we can filter combined things (e.g. size and color)
func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) &&
		a.second.IsSatisfied(p)
}

type BetterFilter struct {
}

func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{name: "apple", color: green, size: small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	fmt.Printf("Green products:\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Printf("Green products(new):\n")
	bf := BetterFilter{}
	greenSpec := ColorSpecification{green}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	// Combinator
	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{largeSpec, greenSpec}
	fmt.Printf("Find large and green products\n")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
