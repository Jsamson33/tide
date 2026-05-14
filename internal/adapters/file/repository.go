package file

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
	"strings"
	"github.com/Jsamson33/tide/internal/domain"
)

//go:embed data/hexagrams/*.md data/trigrams/*.md
var embeddedFS embed.FS

type FileRepository struct{}

func NewFileRepository() *FileRepository {
	return &FileRepository{}
}

func (r *FileRepository) GetHexagramData(h domain.Hexagram) (domain.HexagramData, error) {
	binaryStr := ""
	for _, line := range h.Lines {
		if line == domain.Solid {
			binaryStr += "1"
		} else {
			binaryStr += "0"
		}
	}

	filename := fmt.Sprintf("hexagram_%s.md", binaryStr)
	path := path.Join("data", "hexagrams", filename)

	return r.readFile(path, fmt.Sprintf("Hexagram %s", binaryStr))
}

func (r *FileRepository) GetTrigramData(t domain.Trigram) (domain.HexagramData, error) {
	binaryStr := t.BinaryString()
	filename := fmt.Sprintf("trigram_%s.md", binaryStr)
	path := path.Join("data", "trigrams", filename)

	return r.readFile(path, fmt.Sprintf("Trigram %s", binaryStr))
}

func (r *FileRepository) readFile(path, defaultTitle string) (domain.HexagramData, error) {
	content, err := fs.ReadFile(embeddedFS, path)
	if err != nil {
		return domain.HexagramData{
			Title:       defaultTitle,
			Content:     "Content coming soon...",
			Description: "Data file not found.",
		}, nil
	}

	lines := strings.Split(string(content), "\n")
	data := domain.HexagramData{
		Title:   strings.TrimPrefix(lines[0], "# "),
		Content: string(content),
	}
	if len(lines) > 2 {
		data.Description = lines[2]
	}

	return data, nil
}
