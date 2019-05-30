package main

import (
	"io/ioutil"
	"os"
)

// functions can not have named and unnamed parameters as return type
func ReadFile(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func main() {

}
