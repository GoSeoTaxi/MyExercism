package hamming

import (
	"fmt"
	"strings"
)

func Distance(a, b string) (dif int, err error) {

	ar1 := []rune(a)
	ar2 := []rune(b)

	if (len(ar1)) != (len(ar2)) {
		return 0, fmt.Errorf("different length")
	}

	for index, _ := range ar1 {
		if strings.ToUpper(string(ar1[index])) != strings.ToUpper(string(ar2[index])) {
			dif++
		}
	}
	return dif, err
}
