package main

import (
	"fmt"

	favlib "github.com/gauravssnl/go-recipes/chap01/favoriteslib"
)

func main() {
	fmt.Println("***** Default favorites *****")
	favlib.PrintFavorites()
	favlib.AddFavorite("Kotlin")
	favlib.AddFavorite("Rust")
	favlib.PrintFavorites()
	fmt.Printf("Total number of favorites languages: %d", len(favlib.GetAllFavorites()))
	for k, v := range favlib.GetAllFavorites() {
		fmt.Println("Index:", k, "Value", v)
	}
}
