package utils

import (
	"encoding/base64"
	"net/http"
)

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func RedirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+BasicAuth("username1", "password123"))
	return nil
}
