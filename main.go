package main

import (
	// "encoding/base64"
	"flag"
	"fmt"
	"os"

	"github.com/mervick/aes-everywhere/go/aes256"
)

func main() {

	if len(os.Args) != 5 {
		fmt.Println("encrypt -key \"passwd\" -msg \"message\"")
		os.Exit(1)
	}
	hash := flag.String("key", "", "Decryption Key")
	msg := flag.String("msg", "", "Quoted Message to encrypt")
	flag.Parse()

	if len(*hash) > 0 && len(*msg) > 0 {
		encrypted := aes256.Encrypt(*msg, *hash)
		fmt.Printf("%s\n", encrypted)

	}

}
