package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

type AuthorizationTokenProvider interface {
	GenerateToken() (string, error)
}

type MasterKeyTokenProvider struct {
	Verb         string
	ResourceType string
	ResourceLink string
	Date         string
	Key          string
	KeyType      string
	TokenVersion string
}

func (p MasterKeyTokenProvider) GenerateToken() (token string, err error) {
	decoded, err := base64.StdEncoding.DecodeString(p.Key)

	if err != nil {
		return
	}

	mac := hmac.New(sha256.New, decoded)
	payload := fmt.Sprintf("%v\n%v\n%v\n%v\n\n", strings.ToLower(p.Verb), strings.ToLower(p.ResourceType), p.ResourceLink, strings.ToLower(p.Date))
	mac.Write([]byte(payload))
	hashPayload := mac.Sum(nil)

	sig := base64.StdEncoding.EncodeToString(hashPayload)
	encoded := fmt.Sprintf("type=%v&ver=%v&sig=%v", p.KeyType, p.TokenVersion, sig)
	token = url.QueryEscape(encoded)
	return
}
