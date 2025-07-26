package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}) error {
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func PostJson(url string, target interface{}, data []byte) error {
	bodyReader := bytes.NewReader(data)
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return err
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	return json.NewDecoder(resp.Body).Decode(target)
}
