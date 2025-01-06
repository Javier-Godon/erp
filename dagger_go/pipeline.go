package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"net/http"

	"dagger.io/dagger"
)

func main() {
	checkEnvVariables()

	// Initialize Dagger client
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		log.Fatalf("Failed to connect to Dagger: %v", err)
	}
	defer client.Close()

	username := os.Getenv("USERNAME")
	password := client.SetSecret("password", os.Getenv("CR_PAT"))

	// Define project paths
	projectRoot := "./app"
	source := client.Host().Directory(projectRoot)

	// Create a base container for building and testing the Go app
	appContainer := client.Container().
		From("golang:1.23-slim").
		WithMountedDirectory("/app", source).
		WithWorkdir("/app").
		WithExec([]string{"go", "build", "-o", "app"})

	// Retrieve the latest commit SHA
	repoURL := "https://github.com/Javier-Godon/erp-back"
	latestCommit, err := client.Git(repoURL).
		Branch("main").
		Commit(ctx)
	if err != nil {
		log.Fatalf("Failed to fetch the latest commit: %v", err)
	}
	shortSHA := latestCommit[:7]
	timestamp := time.Now().Format("2006-01-02T15-04")
	imageTag := fmt.Sprintf("%s-%s", shortSHA, timestamp)

	buildArgsMap := map[string]string{
		"tag": latestCommit,
	}
	// Convert map[string]string to []dagger.BuildArg
	buildArgs := []dagger.BuildArg{}
	for key, value := range buildArgsMap {
		buildArgs = append(buildArgs, dagger.BuildArg{
			Name:  key,
			Value: value,
		})
	}

	// Build Docker image
	image := appContainer.
		Directory("/app").
		DockerBuild(dagger.DirectoryDockerBuildOpts{
			Dockerfile: "../Dockerfile", // Path to Dockerfile relative to /app
			BuildArgs:  buildArgs,
		},
		)
	if err != nil {
		log.Fatalf("Failed to build the Docker image: %v", err)
	}
	
	if err != nil {
		log.Fatalf("Failed to build the Docker image: %v", err)
	}

	// Push the image to GHCR
	imageAddress := fmt.Sprintf("ghcr.io/%s/erp-back:%s", username, imageTag)
	_, err = image.WithRegistryAuth("ghcr.io", username, password).
		Publish(ctx, imageAddress)
	if err != nil {
		log.Fatalf("Failed to push the Docker image: %v", err)
	}

	log.Printf("Image published at: %s", imageAddress)

	// Trigger GitHub Action
	dispatchURL := "https://api.github.com/repos/Javier-Godon/cluster-continuous-delivery/dispatches"
	payload := map[string]interface{}{
		"event_type": "image-tag-in-erp-back-dev-updated",
		"client_payload": map[string]string{
			"image_tag": imageTag,
		},
	}
	headers := map[string]string{
		"Authorization": fmt.Sprintf("token %s", os.Getenv("CR_PAT")),
		"Accept":        "application/vnd.github+json",
	}

	// Send the request
	if err := triggerGitHubAction(ctx, dispatchURL, headers, payload); err != nil {
		log.Fatalf("Failed to trigger GitHub Action: %v", err)
	}

	log.Println("GitHub Action triggered successfully")
}

func checkEnvVariables() {
	variablesToCheck := []string{"CR_PAT", "USERNAME"}
	for _, element := range variablesToCheck {
		_, exists := os.LookupEnv(element)
		if !exists {
			//manage this properly
			panic("err")
		}
	}
}

// triggerGitHubAction sends a POST request to GitHub API for repository_dispatch
func triggerGitHubAction(ctx context.Context, url string, headers map[string]string, payload map[string]interface{}) error {
	// Serialize the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers to the request
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusNoContent { // HTTP 204
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
