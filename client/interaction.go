package client

import (
	"fmt"
)

func clearScreen() {
	//msg := "\u001B[2J\n"
	//msg += "\u001B[1;1H"
	//fmt.Print(msg)
}

func askForName(new bool) string {
	clearScreen()
	if new {
		fmt.Print("Welcome, please enter your name: ")
	} else {
		fmt.Print("Enter your name: ")
	}
	var name string
	fmt.Scanln(&name)
	return name
}

func openMenu(options []string) int {
	clearScreen()
	for i, option := range options {
		fmt.Printf("%d. %s\n", i + 1, option)
	}

	for {
		var chosenOption int
		fmt.Scanf("%d", &chosenOption)
		if chosenOption > 0 && chosenOption <= len(options) {
			return chosenOption
		}
	}
}
