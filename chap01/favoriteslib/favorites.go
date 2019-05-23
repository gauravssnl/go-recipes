package favoriteslib

// Store favorites
var favorites []string

// Initialization logic for the package
func init() {
	favorites = make([]string, 3)
	favorites[0] = "Python"
	favorites[1] = "Go"
	favorites[2] = "JavaScript"
}

// AddFavorite - Add a favorite into the in-memory collection (Slice)
func AddFavorite(favorite string) {
	favorites = append(favorites, favorite)
}

// GetAllFavorites - Return all favorites
func GetAllFavorites() []string {
	return favorites
}
