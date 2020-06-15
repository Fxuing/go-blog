package test

import (
	"fmt"
	"testing"
	"time"
)

type Animal struct {
	Name string
}

func (animal *Animal) SayHi() {
	fmt.Println("this name is null")
}

type Dog struct {
	Animal
}

func (dog *Dog) SayHi() {
	fmt.Printf("this name %s\n", dog.Name)
}

type Cat struct {
	Animal
}

func (cat *Cat) SayHi() {
	fmt.Printf("this name %s\n", cat.Name)
}

func TestAnimal(t *testing.T) {
	dog := new(Dog)
	dog.Name = "Dog"
	dog.Animal.SayHi()
	dog.SayHi()
	cat := new(Cat)
	cat.Name = "Cat"
	cat.Animal.SayHi()
	cat.SayHi()
}

// ----------------------------------------------------------------

type Write interface {
	Hello()
}
type JavaOut struct{}

func (j *JavaOut) Hello() {
	fmt.Println("System.out.print(\"hello world\")")
}

type GoOut struct{}

func (g *GoOut) Hello() {
	fmt.Println("fmt.Println(\"hello world\")")
}

func TestInterface(t *testing.T) {
	var jwrite Write = &JavaOut{}
	jwrite.Hello()

	var gwrite Write = &GoOut{}
	gwrite.Hello()
}

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	formate := "2006-01-02 15:04:05"
	t.Log(now.Format(formate))

	timestamp := time.Now().Unix()
	t.Log(timestamp)

}
