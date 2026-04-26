package ascii

import (
	"fmt"
	"os"
	"strings"
)

func getStart(r rune) int {
	return ((int(r) - 32) * 9) + 1
}

func Reader(text string, font string) (string, error) {
	separator := "\n"

	// Build the full path to the font file inside "fonts/" folder
	fontPath := "fonts/" + font

	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		return "", fmt.Errorf("Error reading the file '%s', kindly select another font", font)
	}

	fontStr := strings.ReplaceAll(string(fontBytes), "\r\n", "\n")

	fontSlice := strings.Split(fontStr, separator)
	words := strings.Split(text, "\r\n")
	var result strings.Builder

	for _, w := range words {
		starts := []int{}

		for _, r := range w {
			if r > 126 || r < 32 {
				return "", fmt.Errorf("Invalid character: %c, please use ascii characters only", r)
			}
			starts = append(starts, getStart(r))
		}

		if len(starts) == 0 {
			result.WriteString("\n")
			continue
		}

		for i := 0; i < 8; i++ {
			for _, start := range starts {
				result.WriteString(fontSlice[start+i])
			}
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}
