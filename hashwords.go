package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: hashwords [start number] [end number]")
		os.Exit(1)
	}
	start := new(big.Int)
	if _, ok := start.SetString(os.Args[1], 10); !ok {
		fmt.Println("Please enter a valid start number.")
		os.Exit(1)
	}
	end := new(big.Int)
	if _, ok := end.SetString(os.Args[2], 10); !ok {
		fmt.Println("Please enter a valid end number.")
		os.Exit(1)
	}
	if start.Cmp(end) > 0 {
		fmt.Println("The end number must be greater than the start number.")
		os.Exit(1)
	}
	startTime := time.Now()
	for i := new(big.Int).Set(start); i.Cmp(end) <= 0; i.Add(i, big.NewInt(1)) {
		hash := sha256.Sum256(i.Bytes())
		hashString := base64.StdEncoding.EncodeToString(hash[:])
		elapsed := time.Since(startTime)
		fmt.Printf("%s : %v, elapsed time: %v\n", hashString, i, elapsed)
	}
}
