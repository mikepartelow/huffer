package main

import (
	"fmt"
	"mp/huffer/pkg/huffman"
	"os"
)

func main() {
	const message0 = "For any given degree of noise contamination of a communication channel, it is possible (in theory) to communicate discrete data (digital information) nearly error-free up to a computable maximum rate through the channel."

	encoding, table := huffman.Encode([]rune(message0))

	message1, err := huffman.Decode(encoding, len(message0), table)
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	if message0 != string(message1) {
		fmt.Println("error: message0 != message1")
		os.Exit(1)
	}

	fmt.Println("🤙")
}
