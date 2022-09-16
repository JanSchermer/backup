package util

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

type Uploader struct {
	url    string
	auth   string
	aes    AES
	client http.Client
}

func GetUploader(url string, auth string, aes AES) Uploader {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return Uploader{url, auth, aes, client}
}

func (u Uploader) Upload(file string, hash string, errorChan chan error) {
	req, err := http.NewRequest("POST", u.url, nil)
	if err != nil {
		errorChan <- err
		return
	}
	req.Header.Set("Authorization", u.auth)
	req.Header.Set("File", file)
	req.Header.Set("Hash", hash)
	reader, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		errorChan <- err
		return
	}
	plainBytes, _ := io.ReadAll(reader)
	cipherBytes := u.aes.Encrypt(plainBytes)

	req.Body = io.NopCloser(bytes.NewReader(cipherBytes))
	_, err = u.client.Do(req)
	if err != nil {
		errorChan <- err
	}
}
