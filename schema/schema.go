package schema

// Movie structs
type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director structs
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var Movies []Movie // empty slice
