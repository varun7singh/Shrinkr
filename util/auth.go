package util

import (
	"bytes"
	"net/http"
)

func Get(url string, query string, auth bool) (*http.Response, error) {
	client := http.Client{}
	req, _ := http.NewRequest("GET", url+"?"+query, nil)
	req.Header.Set("Content-Type", "application/json")
	if auth {
		req.Header.Set("Authorization", "Bearer "+GetToken())
	}
	resp, err := client.Do(req)
	return resp, err
}

func Post(url string, body *bytes.Reader, auth bool) (*http.Response, error) {
	client := http.Client{}
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	if auth {
		req.Header.Set("Authorization", "Bearer "+GetToken())
	}
	resp, err := client.Do(req)
	return resp, err
}

func Delete(url string, auth bool) (*http.Response, error) {
	client := http.Client{}
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")
	if auth {
		req.Header.Set("Authorization", "Bearer "+GetToken())
	}
	resp, err := client.Do(req)
	return resp, err
}

func Authenticated() bool {
	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://shrinkr-da1u.onrender.com/shrinkr/user/me", nil)
	req.Header.Set("Authorization", "Bearer "+GetToken())
	resp, _ := client.Do(req)
	if resp.StatusCode == 200 {
		return true
	}
	return false
}
