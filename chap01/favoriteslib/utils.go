package favoriteslib

import "fmt"

// PrintFavorites - Print favorites
func PrintFavorites() {
	for _, v := range favorites { // _ means blank identifier
		fmt.Println(v)
	}
}
