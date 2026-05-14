package domain

// HexagramData represents the content of a hexagram (from markdown files).
type HexagramData struct {
	Title       string
	Content     string
	Description string
}

// HexagramRepository defines the port for accessing hexagram and trigram data.
type HexagramRepository interface {
	GetHexagramData(h Hexagram) (HexagramData, error)
	GetTrigramData(t Trigram) (HexagramData, error)
}
