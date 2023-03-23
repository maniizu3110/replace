package logic

import (
	"fmt"
	"os/exec"
)

const (
	distDir = "./dist/"
)

func Clone(repositoryUrl string, isPrivate bool) error {
	cmd := exec.Command("git", "clone", repositoryUrl)
	cmd.Dir = distDir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git clone error: %v, output: %s", err, string(output))
	}

	return nil
}

func GitAddCommitPush(dir string) error {
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git add error: %v, output: %s", err, string(output))
	}

	cmd = exec.Command("git", "commit", "-m", "auto commit")
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git commit error: %v, output: %s", err, string(output))
	}

	cmd = exec.Command("git", "push")
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git push error: %v, output: %s", err, string(output))
	}

	return nil
}
