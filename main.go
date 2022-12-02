package main

import (
	"exercism.org/luhn"
	"fmt"
)

func main() {
	fmt.Println(luhn.Valid("055b 444 285"))
}
