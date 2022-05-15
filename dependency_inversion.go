package main

import "fmt"

// Dependency Inversion Principle
// HLM should not depend on LLM
// Both should depend on abstractions

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

// Low-level Module (since it is a kind of storage)
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(
	parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child}) // parent is the Parent of child
	r.relations = append(r.relations, Info{child, Child, parent})  // child is a Child of parent

}

// we want to be able to research on the above data
// High-level Module
type Research struct {
	// break DIP
	relationships Relationships //--> this way the LLM depends on HLM, we don't want that
}

func (r *Research) Investigate() {
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" &&
			rel.relationship == Parent {
			fmt.Println("John has a child called ", rel.to.name)
		}
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{relationships}
	r.Investigate()
}
