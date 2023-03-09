package main

import "fmt"

type Writer interface {
	Write(string)
}

type Pen struct {
	Name  string
	Model string
}

func (p Pen) Write(something string) {
	fmt.Printf("Writing with pen: %q\n", something)
}

type Pencil struct {
	Model string
	Color string
}

func (p Pencil) Write(anotherThing string) {
	fmt.Println("Қаламро тез карда истодаам...")
	fmt.Printf("Writing with pencil: %q\n", anotherThing)
}

func main() {
	pen := Pen{
		Model: "Dolphin",
	}

	pencil := Pencil{
		Model: "Architect",
		Color: "Black",
	}

	WriteALetter("Hello dear friend", pen)
	WriteALetter("Hello dear friend", pencil)
}

func WriteALetter(message string, writer Writer) {
	writer.Write(message)
}
