package main

import "fmt"

// Dependency Inversion Principle
// HLM should not depend on LLM
// Both should depend on abstractions

// Typically by LLM we mean what is closer to the hardware/data storage
// and HLM would be the Business Logic stuff.
// Both should depend on abstratctions --> Interfaces in Go

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// abstraction
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// Low-level Module (since it is a kind of storage)
type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Relationships) AddParentAndChild(
	parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child}) // parent is the Parent of child
	r.relations = append(r.relations, Info{child, Child, parent})  // child is a Child of parent

}

// we want to be able to research on the above data
// High-level Module
type Research struct {
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("Jong has chiled called :", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate()
}
