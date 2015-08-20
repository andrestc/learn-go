package main

import (
	"fmt"
)

func say(s string) {
	fmt.Print(s)
}

func main() {
	defer say(" possível")
	say("Viajar é")
	panic("socorro!")
}
