package logic

import (
	"strings"
)

type PathContent struct {
	Path    string
	Content string
}

func ExtractPathsAndContents(input string, dist string) ([]PathContent, error) {
	lines := strings.Split(input, "\n")
	var pathContents []PathContent

	var currentPath string
	var currentContentLines []string

	for _, line := range lines {
		if strings.HasPrefix(line, "PATH:") {
			if currentPath != "" && len(currentContentLines) > 0 {
				pathContents = append(pathContents, PathContent{
					Path:    currentPath,
					Content: strings.Join(currentContentLines, "\n"),
				})
				currentContentLines = nil
			}
			currentPath = strings.TrimSpace(line[5:])
			if !strings.Contains(currentPath, dist) {
				currentPath = dist + currentPath
			}
		} else if strings.HasPrefix(line, "CONTENT:") {
			currentContentLines = append(currentContentLines, strings.TrimSpace(line[8:]))
		} else {
			currentContentLines = append(currentContentLines, line)
		}
	}

	if currentPath != "" && len(currentContentLines) > 0 {
		pathContents = append(pathContents, PathContent{
			Path:    currentPath,
			Content: strings.Join(currentContentLines, "\n"),
		})
	}

	return pathContents, nil
}
