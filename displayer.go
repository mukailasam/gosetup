package main

import "fmt"

func displayer(label string, deps ...string) {
	fmt.Println("\n" + blueText(label))
	for k, v := range deps {
		fmt.Printf(yellowText("[%d]"), k+1)
		fmt.Printf(greenText(" %s\n"), v)
	}
}
