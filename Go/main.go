package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

const API_URL = "https://fontys.cloud-builders.nl/api/v1/send-image"

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("Environment variable API_KEY is not set")
		os.Exit(1)
	}

	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <x> <y> <image_path>")
		os.Exit(1)
	}

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("x must be an integer")
		os.Exit(1)
	}
	y, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("y must be an integer")
		os.Exit(1)
	}

	imagePath := os.Args[3]
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Prepare the request
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	part, err := writer.CreateFormFile("image", imagePath)
	if err != nil {
		fmt.Println("Error creating form file:", err)
		os.Exit(1)
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	_, err = part.Write(fileBytes)
	if err != nil {
		fmt.Println("Error writing file bytes:", err)
		os.Exit(1)
	}
	writer.Close()

	url := fmt.Sprintf("%s?x=%d&y=%d", API_URL, x, y)
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}
	req.Header.Set("X-Api-Key", apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}
