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
		[]byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

// How to deal with SRP break.

// To bypass this issue for save we could create a free standing function
var LineSeperator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.enties, LineSeperator)), 0644)
}

// --------------------------------------
// Another way would be to turn persistence into a struct
type Persistence struct {
	lineSeperator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.enties, p.lineSeperator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("helloo")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String())

	//
	SaveToFile(&j, "journal.txt")

	//
	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
