package main

import (
	"encoding/json"
	"bufio"
	"log"
	"os"

	"github.com/telia-oss/github-pr-resource"
)

func main() {
	var request resource.CheckRequest

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	if err := json.Unmarshal([]byte(text), &request); err != nil {
		log.Fatalf("failed to unmarshal request: %s -> %s", err, text)
	}

	if err := request.Source.Validate(); err != nil {
		log.Fatalf("invalid source configuration: %s", err)
	}
	github, err := resource.NewGithubClient(&request.Source)
	if err != nil {
		log.Fatalf("failed to create github manager: %s", err)
	}
	response, err := resource.Check(request, github)
	if err != nil {
		log.Fatalf("check failed: %s", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		log.Fatalf("failed to marshal response: %s", err)
	}
}
