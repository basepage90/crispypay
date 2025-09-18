package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// 파트너사와 HTTP POST 통신을 위함 유틸
func PostJSON[T any](client *http.Client, url string, request any, response *T) error {
	// to json
	jsonData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("JSON 직렬화 실패: %w", err)
	}

	// make request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("요청 생성 실패: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer partner-api-key")

	// call
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP 요청 실패: %w", err)
	}

	defer resp.Body.Close()

	// check status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP 오류 (%d): %s", resp.StatusCode, string(body))
	}

	// bind response
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return fmt.Errorf("response 디코딩 실패: %w", err)
	}

	return nil
}

// 파트너사와 HTTP GET 통신을 위한 유틸 - 미사용
func Get[T any](client *http.Client, url string, response *T) error {
	// make request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("요청 생성 실패: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer partner-api-key")

	// call
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP 요청 실패: %w", err)
	}
	defer resp.Body.Close()

	// check status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP 오류 (%d): %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		if errors.Is(err, io.EOF) {
			var zero T
			*response = zero
			return nil
		}
		return fmt.Errorf("response 디코딩 실패: %w", err)
	}

	return nil
}
