//go:build ignore

package main

import "fmt"

var TokenTypes = map[string]int{
	"View":      0,
	"StartTag":  1,
	"EndTag":    2,
	"Comment":   3,
	"Text":      4,
	"Code":      5,
	"EndOfFile": 6,
}

var TokenModifiers = map[string]int{
	"Component": 0,
}

func generateArrayFromMap(m map[string]int) []string {
	var tmp = make(map[int]string)
	for k, v := range m {
		tmp[v] = k
	}
	var a = []string{}
	for i := 0; i < len(tmp); i++ {
		a = append(a, tmp[i])
	}
	return a
}

func main() {

	fmt.Printf("tokenTypes: %v\n", generateArrayFromMap(TokenTypes))
	fmt.Printf("tokenModifiers: %v\n", generateArrayFromMap(TokenModifiers))

}
