package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

// A Type should have one primary responsibility
// and as a result should have one reason to change

var entryCount = 0

type Journal struct {
	enties []string
}

func (j *Journal) String() string {
	return strings.Join(j.enties, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.enties = append(j.enties, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

// --- Break single responsibility principle ---
// seperation of concerns

// From here we BREAK single responsibiblity principle,
// because the responsibility of Journal struct is not to deal with Percistence.
// Persistence can be handled by a seperate component (seperate package/struct)

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()))
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

func main() {

}
