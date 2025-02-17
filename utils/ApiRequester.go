package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func PVEAPIRequester(address, path, auth, method, payload string) (string, error) {
	url := fmt.Sprintf("https://%s:8006/%s", address, path)

	var reqBody io.Reader
	if payload != "" && (method == "POST" || method == "PUT" || method == "PATCH") {
		reqBody = bytes.NewBuffer([]byte(payload))
	} else {
		reqBody = nil
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return "Error", err
	}
	req.Header.Set("Authorization", auth)

	if payload != "" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return "Error", err
	}
	defer resp.Body.Close()

	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return "Error", err
	}

	return buf.String(), nil
}
