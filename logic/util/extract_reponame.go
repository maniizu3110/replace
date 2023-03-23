package util

import (
	"fmt"
	"strings"
)

func ExtractRepoName(repoUrl string) (string, error) {
	if !strings.Contains(repoUrl, ".git") {
		return "", fmt.Errorf("Invalid repository URL: %s", repoUrl)
	}

	parts := strings.Split(repoUrl, "/")
	if len(parts) < 1 {
		return "", fmt.Errorf("Invalid repository URL: %s", repoUrl)
	}

	repoNameWithGit := parts[len(parts)-1]

	repoName := strings.TrimSuffix(repoNameWithGit, ".git")
	return repoName, nil
}
