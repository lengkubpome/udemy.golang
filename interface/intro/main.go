package main

import "fmt"

type Sword struct {
	name string
}

func (s Sword) detail() string {
	return "Super cold Iced sword"
}

func (s Sword) cost() int {
	return 100
}

type Bow struct {
	name string
}

func (b Bow) detail() string {
	return "Magic fire-bow"
}
func (b Bow) cost() int {
	return 80
}

type Item interface {
	cost() int
}

type Weapon interface {
	detail() string
	Item
}

func attack(w Weapon) {
	fmt.Printf("attack by : %s. Item cost : %d\n", w.detail(), w.cost())
}

func main() {
	sword := Sword{name: "Ice-sword"}
	bow := Bow{name: "Fier-Bow"}

	var w Weapon
	w = sword
	attack(w)
	w = bow
	attack(w)
}
