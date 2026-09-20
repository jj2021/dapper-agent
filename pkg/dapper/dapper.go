package dapper

import (
	"bufio"
	"bytes"
	"encoding/json"
    "fmt"
    "net/http"
	"os"
)

type ollamaResponse struct {
	Model string `json:"model"`
	CreatedAt string `json:"created_at"`
	Message *ollamaResponseMsg `json:"message"`
	DoneReason string `json:"done_reason"`
	Done bool `json:"done"`
	TotalDuration int `json:"total_duration"`
	LoadDuration int `json:"load_duration"`
	PromptEvalCount int `json:"prompt_eval_count"`
	PromptEvalDuration int `json:"prompt_eval_duration"`
	EvalCount int `json:"eval_count"`
	EvalDuration int `json:"eval_duration`
}

type ollamaResponseMsg struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

func Run() {
	fmt.Printf("G'day... How may I be of assistance?\n\n")

	for {
		fmt.Printf("> ")
		// get user prompt
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if scanner.Text() == "/quit" {
			break
		}

		// form request body
		reqBody := fmt.Sprintf(`{"model": "granite4.2:3b", "messages": [{"role": "user", "content": "%s"}]}`, scanner.Text())

		// make ollama request
		client := &http.Client{}
		reqBytes := make([]byte, 1024)
		reqBytes = []byte(reqBody)
		reqBytesBuf := bytes.NewBuffer(reqBytes)
		req, err := http.NewRequest("POST", "http://127.0.0.1:11434/api/chat", reqBytesBuf)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error making chat request: %s\n", err.Error())
			return
		}

		// process the ollama response
		respBytes := make([]byte, 1024)
		var respErr error = nil
		ollamaResp := &ollamaResponse{Message:&ollamaResponseMsg{}}
		for {
			if respErr != nil {
				break
			}
			n, err := resp.Body.Read(respBytes)
			respErr = err

			unmarshalErr := json.Unmarshal(respBytes[:n], ollamaResp)
			if unmarshalErr != nil {
				fmt.Printf("%s\n", unmarshalErr.Error())
				break
			}

			fmt.Printf("%s", ollamaResp.Message.Content)

			if ollamaResp.Done {
				fmt.Printf("\n\n")
				break
			}
		}
	}
}
