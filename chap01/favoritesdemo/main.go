package main

import (
	"fmt"

	"github.com/gauravssnl/go-recipes/chap01/favoriteslib"
)

func main() {
	fmt.Println("***** Default favorites *****")
	favoriteslib.PrintFavorites()
	favoriteslib.AddFavorite("Kotlin")
	favoriteslib.AddFavorite("Rust")
	favoriteslib.PrintFavorites()
	fmt.Printf("Total number of favorites languages: %d", len(favoriteslib.GetAllFavorites()))
	for k, v := range favoriteslib.GetAllFavorites() {
		fmt.Println("Index:", k, "Value", v)
	}
}
