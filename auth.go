package documentdb

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

// AuthorizationTokenProvider defines the interface needed for creating Cosmos DB authentication tokens
type AuthorizationTokenProvider interface {
	GenerateToken() (string, error)
}

// MasterKeyTokenProvider provides a type for creating master key authorization tokens
type MasterKeyTokenProvider struct {
	Verb         string
	ResourceType string
	ResourceLink string
	Date         string
	Key          string
	KeyType      string
	TokenVersion string
}

// GenerateToken generates the master key authorization token for authenticating with the Cosmos DB REST API
func (p *MasterKeyTokenProvider) GenerateToken() (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(p.Key)

	if err != nil {
		return "", err
	}

	mac := hmac.New(sha256.New, decoded)
	payload := fmt.Sprintf("%v\n%v\n%v\n%v\n\n", strings.ToLower(p.Verb), strings.ToLower(p.ResourceType), p.ResourceLink, strings.ToLower(p.Date))
	mac.Write([]byte(payload))
	hashPayload := mac.Sum(nil)

	sig := base64.StdEncoding.EncodeToString(hashPayload)
	encoded := fmt.Sprintf("type=%v&ver=%v&sig=%v", p.KeyType, p.TokenVersion, sig)

	return url.QueryEscape(encoded), nil
}
