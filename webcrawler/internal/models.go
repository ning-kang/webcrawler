package internal

type Book struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Price        float32 `json:"price"`
	Tax          float32 `json:"tax"`
	Availability string  `json:"availability"`
	Rating       int     `json:"rating"`
	Description  string  `json:"description"`
	UPC          string  `json:"upc"`
	Category     string  `json:"category"`
}
