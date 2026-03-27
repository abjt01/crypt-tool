package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("crypt-tool")
}
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		printHelp()
		os.Exit(0)
	}

	action := os.Args[1]

	switch action {
	case "help":
		printHelp()
	default:
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println(`crypt-tool — text & file encryption

Usage:
  go run . encrypt text "YOUR TEXT" <key>
  go run . decrypt text "ENCRYPTED"  <key>
  go run . encrypt file <path>
  go run . decrypt file <path>

Examples:
  go run . encrypt text "HELLO WORLD" 5
  go run . decrypt text "MJQQT BTWQI" 5
  go run . encrypt file test.jpeg
  go run . decrypt file test.jpeg`)
}