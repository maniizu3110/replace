package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const systemMessage = "You are an automatic code generator. Your response to the user must follow the following format: PATH: Path to file CONTENT:Contents of file. Output other than the path to the file and the contents of the file is prohibited."

type ChatCompletion struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func CallGPT4(messages []string) ([]string, error) {
	url := "https://api.openai.com/v1/chat/completions"
	method := "POST"

	userMessages := make([]map[string]string, len(messages))
	for i, msg := range messages {
		userMessages[i] = map[string]string{
			"role":    "user",
			"content": msg,
		}
	}
	userMessages = append([]map[string]string{{"role": "system", "content": systemMessage}}, userMessages...)

	fmt.Println(userMessages)
	requestPayload := map[string]interface{}{
		"model":             "gpt-4",
		"messages":          userMessages,
		"max_tokens":        2500,
		"top_p":             0.8,
		"n":                 1,
		"frequency_penalty": 0.0,
		"presence_penalty":  0.6,
	}

	jsonPayload, err := json.Marshal(requestPayload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}

	payload := strings.NewReader(string(jsonPayload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	key := os.Getenv("OPEN_API_KEY")
	if key == "" {
		fmt.Println("OPEN_API_KEY is not set")
		return nil, err
	}
	authorization := "Bearer " + key
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorization)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var cc ChatCompletion
	err = json.Unmarshal(body, &cc)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	contents := make([]string, len(cc.Choices))
	for _, choice := range cc.Choices {
		contents = append(contents, choice.Message.Content)
	}

	return contents, nil
}
