package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type resp struct {
	Url string `json:"url"`
}

func Upload(imageData []byte, filename string) (string, error) {
	// TODO: upload to the python service
	img := bytes.NewBuffer(imageData)
	var reqBody bytes.Buffer
	writer := multipart.NewWriter(&reqBody)
	fileField, err := writer.CreateFormFile("image", "image")
	io.Copy(fileField, img)
	if err != nil {
		return "", err
	}
	err = writer.WriteField("toRemove", "false")
	if err != nil {
		return "", err
	}
	err = writer.WriteField("filename", filename)
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}
	uploadURL := os.Getenv("UPLOAD_URL")
	if uploadURL == "" {
		log.Fatalf("there is no UPLOAD_URL environment variable")
	}

	req, err := http.NewRequest(http.MethodPost, uploadURL, &reqBody)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	var res resp
	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		return "", err
	}
	log.Printf("response: %v", string(res.Url))

	return res.Url, nil
}
