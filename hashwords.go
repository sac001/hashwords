package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"math/big"
	"strconv"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func uintGen(min, max uint64) {
	for i := min; ; i++ {
		h := strconv.FormatUint(i, 16)
		// strconv should be faster than fmt

		for {
			if len(h)%2 == 0 {
				break
			}
			h = "0" + h
		}
		// hex package only works in multiples of 2

		dst, err := hex.DecodeString(h)
		if err != nil {
			panic(err)
		}

		hash := sha256.Sum256([]byte(dst))
		fmt.Println(Reverse(base64.StdEncoding.EncodeToString(hash[:])))

		if i >= max {
			break
		}
	}
}

func bigIntGen(min, max *big.Int) {

	for i := min; i.Cmp(max) < 0; i.Add(i, big.NewInt(1)) {
		h := i.Text(16)

		for {
			if len(h)%2 == 0 {
				break
			}
			h = "0" + h
		}
		// hex package only works in multiples of 2

		dst, err := hex.DecodeString(h)
		if err != nil {
			panic(err)
		}

		hash := sha256.Sum256([]byte(dst))
		fmt.Println(Reverse(base64.StdEncoding.EncodeToString(hash[:])))

	}
}

func main() {
	bigint := flag.Bool("bigint", false, "use big int implementaion?")
	max := flag.String("max", "300", "max. number to generate")
	min := flag.String("min", "1", "min. number to generate")
	flag.Parse()

	if *bigint {
		m := &big.Int{}
		mi := &big.Int{}
		if _, success := m.SetString(*max, 10); !success {
			panic("unable to convert max to big int")
		}
		m.Add(m, big.NewInt(1))

		if _, success := mi.SetString(*min, 10); !success {
			panic("unable to convert max to big int")
		}
		bigIntGen(mi, m)
	} else {
		m, err := strconv.ParseUint(*max, 10, 64)
		if err != nil {
			panic(err)
		}
		mi, err := strconv.ParseUint(*min, 10, 64)
		if err != nil {
			panic(err)
		}

		uintGen(mi, m)
	}
}
