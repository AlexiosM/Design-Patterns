package main

type Document struct {
}

type Machine interface {
	Print(Document)
	Fax(Document)
	Scan(Document)
}

type MultiFunctionPrinter struct {
}

func (m *MultiFunctionPrinter) Print(d Document) {
}
func (m *MultiFunctionPrinter) Fax(d Document) {
}
func (m *MultiFunctionPrinter) Scan(d Document) {
}

type OldFashionPrinter struct {
}

func (o *OldFashionPrinter) Print(d Document) {
	//ok
}

// Deprecated: ...
func (o *OldFashionPrinter) Fax(d Document) {
	panic("Operation not supported")
}

// Deprecated: ...
func (o *OldFashionPrinter) Scan(d Document) {
	panic("Operation not supported")
}

// ISP (Interface Segregation Principle)
// try to breakup an interface into seperate parts that people definitely need
// (there is no guarantee that if someone needs printing, also needs faxing)
// Let's split Print and Scan into diferrent interfaces

type Printer interface {
	Print(d Document)
}
type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (mp *MyPrinter) Print(d Document) {

}

type PhotoCopier struct {
}

func (p *PhotoCopier) Print(d Document) {
}

func (p *PhotoCopier) Scan(d Document) {
}

// Combining interfaces
type MultiFunctionDevice interface {
	Printer
	Scanner
	//Fax
}

// If we want to build a MultifunctionMachine
// and there are already implemented a Printed and a Scanner implemented as seperate components
// Then we can use a Decorator Design Pattern

// decorator
type MultifunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (mfm MultifunctionMachine) Print(d Document) {
	mfm.printer.Print(d)
}
func (mfm MultifunctionMachine) Scan(d Document) {
	mfm.scanner.Scan(d)
}

func main() {
}
