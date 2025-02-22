package internal

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	fp "path/filepath"
	"strings"
)

var api_host = "https://api.tempfish.com/api/v1/"

func UploadFiles(files []string) (*bytes.Buffer, error) {
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

	resp, err := http.Post(api_host+"push", writer.FormDataContentType(), body)
	if err != nil {
		return nil, fmt.Errorf("error making POST request: %w", err)
	}
	defer resp.Body.Close()

	respBody := &bytes.Buffer{}
	_, err = io.Copy(respBody, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading server response: %w", err)
	}
	fmt.Println("Server response body:", respBody.String())

	return body, nil
}

func GetFiles(dst string, directory string) error {
	directoryName := fp.Join(dst, "/tfs-files")
	res, err := http.Get(api_host + "pull/" + directory)

	if err != nil {
		fmt.Printf("error making request: %v\n", err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("no files found: %v", res.Status)
	}

	contentType := res.Header.Get("Content-Type")

	if !strings.HasPrefix(contentType, "multipart/") {
		fmt.Println("response is not multipart")
		return fmt.Errorf("response is not multipart")
	}

	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		fmt.Printf("error parsing media type: %v\n", err)
		return err
	}

	boundary, ok := params["boundary"]
	if !ok {
		fmt.Println("no boundary found in Content-Type")
		return fmt.Errorf("no boundary found in Content-Type")
	}

	mr := multipart.NewReader(res.Body, boundary)

	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error getting next part: %v\n", err)
			return err
		}
		if part == nil {
			fmt.Println("error: part is nil")
			return fmt.Errorf("part is nil")
		}

		fileName := part.FileName()
		fmt.Printf("Processing part with filename: %s\n", fileName)
		if fileName == "" {
			continue
		}

		if _, err := os.Stat(directoryName); os.IsNotExist(err) {
			err = os.Mkdir(directoryName, os.ModePerm)
			if err != nil {
				panic("Unable to create uploads directory")
			}
		}

		out, err := os.Create(fp.Join(directoryName, fileName))
		if err != nil {
			fmt.Printf("error creating file %s: %v\n", fileName, err)
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, part)
		if err != nil {
			fmt.Printf("error saving file %s: %v\n", fileName, err)
			return err
		}
	}

	fmt.Println("Files downloaded successfully")
	return nil
}
