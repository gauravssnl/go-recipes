package main

import (
	"fmt"

	"github.com/gauravssnl/go-recipes/chap01/strutils"
)

func main() {
	s1, s2 := "python", "PYTHON"
	fmt.Println(strutils.ToUpperCase(s1))
	fmt.Println(strutils.ToLowerCase(s2))
	fmt.Println(strutils.FirstToupper(s1))
}
