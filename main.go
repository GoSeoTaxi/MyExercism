package main

import (
	"exercism.org/clock"
	"exercism.org/isogram"
	"fmt"
)

func main() {

	fmt.Println(clock.New(23, 59).Add(2))

	fmt.Println(isogram.IsIsogram("АБВГДа"))

}
