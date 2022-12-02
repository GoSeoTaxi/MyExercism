package luhn

import (
	"strconv"
)

func Valid(id string) bool {
	cleanID := ""
	for i := 0; i < len(id); i++ {
		if id[i] == 32 {
			continue
		}
		cleanID = cleanID + string(id[i])
	}

	tempString := ""
	counter := 0
	for i := len(cleanID) - 1; i >= 0; i-- {
		counter++
		item, err := strconv.ParseInt(string(cleanID[i]), 10, 8)
		if err != nil {
			return false
		}
		if (counter % 2) == 0 {
			item = item * 2
			if item > 9 {
				item = item - 9
			}
		}
		tempString = tempString + strconv.Itoa(int(item))
	}

	count := int64(0)
	for i := 0; i < len(tempString); i++ {
		i, _ := strconv.ParseInt(string(tempString[i]), 10, 8)
		count = count + i
	}
	if count%10 != 0 || len(cleanID) <= 1 {
		return false
	}
	return true

}
