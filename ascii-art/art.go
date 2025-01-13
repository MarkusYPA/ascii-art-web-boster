package asciiart

import (
	"fmt"
	"os"
)

func squeeze(art []string) []string {
	newLines := []string{}
	for _, line := range art {
		if line != "" {
			newLines = append(newLines, line)
		}
	}
	return newLines
}

func removeCarReturns(art []byte) []byte {
	result := []byte{}
	for _, char := range art {
		if char != '\r' {
			result = append(result, char)
		}
	}
	return result
}

func toLines(art []byte) []string {
	art = removeCarReturns(art)

	lines := []string{}

	// Write art contents to slice of strings
	currentLine := ""
	for _, char := range art {
		if char != '\n' {
			currentLine += string(char)
		} else {
			lines = append(lines, currentLine)
			currentLine = ""
		}
	}

	// Remove empty lines
	lines = squeeze(lines)

	return lines
}

func removeOneIfAllEmpty(inputLines []string) []string {
	allEmpty := true

	for _, line := range inputLines {
		if line != "" {
			allEmpty = false
			break
		}
	}

	if allEmpty {
		inputLines = inputLines[:len(inputLines)-1] // remove last element
	}

	return inputLines
}

func artLine(inLine string, bannerLines []string) string {
	result := ""
	for i := 0; i < 8; i++ {
		for _, char := range inLine {
			bannerLineIndex := int(char-' ')*8 + i
			result += bannerLines[bannerLineIndex]
		}
		result += "\n"
	}

	return result
}

func CreateArt(inputLines []string, style string) (string, error) {

	bannerFileContent, err := os.ReadFile("./banners/" + style + ".txt")
	if err != nil {
		fmt.Println("Error reading banner file")
		return "", err
	}

	bannerLines := toLines(bannerFileContent)
	inputLines = removeOneIfAllEmpty(inputLines)

	art := ""
	for _, inLine := range inputLines {
		if inLine == "" {
			art += "\n"
		} else {
			art += artLine(inLine, bannerLines)
		}
	}

	return art, nil

}
