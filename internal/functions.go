package internal

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func uploadFiles(files []string) (*bytes.Buffer, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for _, filePath := range files[1:] { //Ignoring the first arg
		file, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("error opening file %s: %w", filePath, err)
		}
		defer file.Close()

		part, err := writer.CreateFormFile("files", file.Name())
		if err != nil {
			return nil, fmt.Errorf("error creating form file for %s: %w", filePath, err)
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return nil, fmt.Errorf("error copying file %s: %w", filePath, err)
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing writer: %w", err)
	}

	resp, err := http.Post("http://localhost:8080/api/v1/push", writer.FormDataContentType(), body)
	if err != nil {
		return nil, fmt.Errorf("error making POST request: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println("Server response status:", resp.Status)
	respBody := &bytes.Buffer{}
	_, err = io.Copy(respBody, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading server response: %w", err)
	}
	fmt.Println("Server response body:", respBody.String())

	return body, nil
}
