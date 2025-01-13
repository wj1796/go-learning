package main

import (
	"fmt"
)

func main() {

	text := "CSOITEUIWUIZNSROCNKFD"
	KEY := "GOLANG"
	newmessage := ""
	keyindex := 0

	for _, c := range text {
		k := rune(KEY[keyindex]) - 'A'
		c = c - 'A'
		ans := (c+26-k)%26 + 'A'
		newmessage += string(ans)
		keyindex++
		keyindex = keyindex % 6
	}
	fmt.Println(newmessage)

}
