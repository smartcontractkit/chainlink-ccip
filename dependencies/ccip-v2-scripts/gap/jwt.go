package gap

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// OIDCResponse represents the expected response structure
type OIDCResponse struct {
	Value string `json:"value"`
}

func decodeBase64URL(encodedToken string) (string, error) {
	decodedString, err := base64.RawURLEncoding.DecodeString(encodedToken)
	if err != nil {
		return "", fmt.Errorf("unable to decode token: %v", err)
	}
	return string(decodedString), nil
}

func CheckToken(validToken string) (map[string]interface{}, map[string]interface{}, error) {
	// Split the token into parts
	parts := strings.Split(validToken, ".")
	if len(parts) != 3 {
		logger.Fatalf("Invalid JWT: must have three parts")
	}

	// Decode the header
	headerJSON, err := decodeBase64URL(parts[0])
	if err != nil {
		logger.Fatalf("Error decoding header: %v", err)
	}

	// Decode the payload
	payloadJSON, err := decodeBase64URL(parts[1])
	if err != nil {
		logger.Fatalf("Error decoding payload: %v", err)
	}

	// Parse the header as JSON
	var header map[string]interface{}
	if err := json.Unmarshal([]byte(headerJSON), &header); err != nil {
		logger.Fatalf("Error parsing header JSON: %v", err)
	}

	// Parse the payload as JSON
	var payload map[string]interface{}
	if err := json.Unmarshal([]byte(payloadJSON), &payload); err != nil {
		logger.Fatalf("Error parsing payload JSON: %v", err)
	}

	// Print the decoded header and payload
	// fmt.Println("Header:")
	// prettyPrintJSON(header)

	// fmt.Println("Payload:")
	// prettyPrintJSON(payload)
	// todo: we can add some validation here to ensure that the token token is valid

	return header, payload, nil
}

func FetchJWTTokenForGAP(ctx context.Context) (string, error) {
	// Fetch environment variables required for OIDC request
	oidcURL := os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
	oidcToken := os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")

	if oidcURL == "" || oidcToken == "" {
		return "", fmt.Errorf("ACTIONS_ID_TOKEN_REQUEST_URL or ACTIONS_ID_TOKEN_REQUEST_TOKEN is not set")
	}

	// Append audience parameter to URL (optional but recommended)
	audience := "gap"
	fullURL := fmt.Sprintf("%s&audience=%s", oidcURL, audience)

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Set Authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oidcToken))

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error fetching OIDC token: %w", err)
	}
	defer resp.Body.Close()

	// Parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var oidcResp OIDCResponse
	if err := json.Unmarshal(body, &oidcResp); err != nil {
		return "", fmt.Errorf("error parsing JSON response: %w", err)
	}
	encodedToken := oidcResp.Value
	if encodedToken == "" {
		return "", fmt.Errorf("oidc response is empty")
	}

	_, _, err = CheckToken(encodedToken)
	if err != nil {
		return "", fmt.Errorf("invalid token: %s", err)
	}

	return encodedToken, nil
}
