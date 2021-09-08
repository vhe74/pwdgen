package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const DEFAULTLEN = 10

func main() {
	args := os.Args
	p := ""
	if len(args) > 1 {
		l, err := strconv.Atoi(args[1])
		if err != nil {
			p = RandStringBytes(DEFAULTLEN)
		}
		p = RandStringBytes(l)
	} else {
		p = RandStringBytes(DEFAULTLEN)
	}
	fmt.Println(p)
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ123456789?./+,;:=@&(!)"
const numbers = "123456789"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		rand.Seed(time.Now().UTC().UnixNano())
		b[i] = letters[rand.Intn(len(letters))]
	}
	//Force the presence of 1 number
	b[n-2] = numbers[rand.Intn(len(numbers))]
	return string(b)
}
