package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"crypt-tool/filecrypt"

	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 3 {
		printHelp()
		os.Exit(0)
	}

	action := os.Args[1]
	mode := os.Args[2]

	switch action {
	case "help":
		printHelp()

	case "encrypt":
		switch mode {
		case "text":
			handleTextEncrypt()
		case "file":
			handleFileEncrypt()
		default:
			fmt.Println("Mode must be 'text' or 'file'.")
			os.Exit(1)
		}

	case "decrypt":
		switch mode {
		case "text":
			handleTextDecrypt()
		case "file":
			handleFileDecrypt()
		default:
			fmt.Println("Mode must be 'text' or 'file'.")
			os.Exit(1)
		}

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

// text(simple letter shift cipher)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func shiftAlphabet(key int) string {
	n := len(alphabet)
	k := key % n
	return alphabet[n-k:] + alphabet[:n-k]
}

func encryptText(key int, plain string) string {
	shifted := shiftAlphabet(key)
	var sb strings.Builder
	for _, r := range strings.ToUpper(plain) {
		pos := strings.IndexRune(alphabet, r)
		if pos >= 0 {
			sb.WriteByte(shifted[pos])
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func decryptText(key int, enc string) string {
	shifted := shiftAlphabet(key)
	var sb strings.Builder
	for _, r := range strings.ToUpper(enc) {
		pos := strings.IndexRune(shifted, r)
		if pos >= 0 {
			sb.WriteByte(alphabet[pos])
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func handleTextEncrypt() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: go run . encrypt text \"TEXT\" <key>")
		os.Exit(1)
	}
	key, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("Key must be an integer.")
		os.Exit(1)
	}
	fmt.Println("Encrypted:", encryptText(key, os.Args[3]))
}

func handleTextDecrypt() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: go run . decrypt text \"TEXT\" <key>")
		os.Exit(1)
	}
	key, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("Key must be an integer.")
		os.Exit(1)
	}
	fmt.Println("Decrypted:", decryptText(key, os.Args[3]))
}

func handleFileEncrypt() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run . encrypt file <path>")
		os.Exit(1)
	}
	f := os.Args[3]
	filecrypt.Encrypt(f, []byte("password"))
	fmt.Println("File encrypted successfully.")
}

func handleFileDecrypt() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run . decrypt file <path>")
		os.Exit(1)
	}
	f := os.Args[3]
	filecrypt.Decrypt(f, []byte("password"))
	fmt.Println("File decrypted successfully.")
}

func getPassword() []byte {
	fmt.Print("Enter password: ")
	p1, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Print("\nConfirm password: ")
	p2, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if !bytes.Equal(p1, p2) {
		fmt.Println("Passwords don't match, try again.")
		return getPassword()
	}
	return p1
}
