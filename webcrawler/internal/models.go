package internal

type Book struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Price        float32 `json:"price"`
	Availability int     `json:"availability"`
	Rating       int     `json:"rating"`
	Description  string  `json:"description"`
	UPC          string  `json:"upc"`
	Category     string  `json:"category"`
}
