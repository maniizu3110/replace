package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/maniizu3110/replace/logic"
	"github.com/maniizu3110/replace/logic/util"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

const (
	repositoryUrl = "https://github.com/maniizu3110/gittest.git"
)

func main() {
	repositoryURL := os.Getenv("GITHUB_REPOSITORY")
	if repositoryURL == "" {
		panic("GITHUB_REPOSITORY is not set")
	}
	message := os.Getenv("MESSAGE_TO_GPT")
	if message == "" {
		panic("MESSAGE_TO_GPT is not set")
	}
	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("Clone repository...")
	go func() {
		defer wg.Done()
		logic.Clone(repositoryUrl, false)
	}()
	fmt.Println("Clone completed!")

	wg.Wait()

	repoName, err := util.ExtractRepoName(repositoryUrl)
	if err != nil || repoName == "" {
		panic(err)
	}
	dist := "dist/" + repoName + "/"
	repoContent, err := logic.DirToStrings(dist)
	if err != nil {
		panic(err)
	}
	repoContent = append(repoContent, message)

	var (
		gptResponses []string
	)
	logrus.Info("Call GPT-4...")
	for _, gptRequest := range repoContent {
		wg.Add(1)
		go func(request string) {
			defer wg.Done()
			gptResponses, err = logic.CallGPT4(repoContent)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}(gptRequest)
	}
	wg.Wait()
	logrus.Info("Call GPT-4 completed!")

	for _, gptResponse := range gptResponses {
		pathContents, err := logic.ExtractPathsAndContents(gptResponse, dist)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for _, pc := range pathContents {
			fmt.Println(pc.Path)
			logic.WiteFile(pc.Path, pc.Content)
		}
	}

	logrus.Info("Git add, commit, push...")
	err = logic.GitAddCommitPush(dist)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	logrus.Info("Git add, commit, push completed!")
	err = os.RemoveAll("dist")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
