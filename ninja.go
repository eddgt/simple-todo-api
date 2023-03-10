// package main

import "fmt"

func main() {
	fmt.Println("Hello Ninja!")

	type ninja struct {
		name    string
		weapons []string
		level   int
	}

	shinobi := ninja{name: "Shinobi"}
	shinobi = ninja{
		name:    "Shinobi",
		weapons: []string{"Ninja Star"},
		level:   1,
	}

	fmt.Println(shinobi)
}
