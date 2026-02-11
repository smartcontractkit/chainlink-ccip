package jd

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	ciptypes "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"golang.org/x/oauth2"
)

// CognitoTokenSource implements oauth2.TokenSource using AWS Cognito
// USER_PASSWORD_AUTH flow. It automatically refreshes tokens using
// the REFRESH_TOKEN_AUTH flow when the access token expires.
type CognitoTokenSource struct {
	client       *cip.Client
	clientID     string
	clientSecret string
	username     string
	password     string

	mu           sync.Mutex
	accessToken  string
	refreshToken string
	expiry       time.Time
}

// NewCognitoTokenSource creates a new Cognito token source.
// It does NOT perform the initial authentication; that happens on the first
// call to Token().
func NewCognitoTokenSource(ctx context.Context, clientID, clientSecret, username, password, region string) (*CognitoTokenSource, error) {
	cfg, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("cognito: failed to load AWS config: %w", err)
	}
	return &CognitoTokenSource{
		client:       cip.NewFromConfig(cfg),
		clientID:     clientID,
		clientSecret: clientSecret,
		username:     username,
		password:     password,
	}, nil
}

// Token returns a valid OAuth2 access token. If the current token is expired
// or absent, it obtains a new one via Cognito.
func (c *CognitoTokenSource) Token() (*oauth2.Token, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Return cached token if still valid (with 30s buffer).
	if c.accessToken != "" && time.Now().Before(c.expiry.Add(-30*time.Second)) {
		return &oauth2.Token{
			AccessToken: c.accessToken,
			Expiry:      c.expiry,
		}, nil
	}

	// Try refresh if we have a refresh token.
	if c.refreshToken != "" {
		tok, err := c.doRefresh()
		if err == nil {
			return tok, nil
		}
		// Fall through to full auth on refresh failure.
	}

	// Full USER_PASSWORD_AUTH flow.
	return c.doPasswordAuth()
}

func (c *CognitoTokenSource) doPasswordAuth() (*oauth2.Token, error) {
	params := map[string]string{
		"USERNAME": c.username,
		"PASSWORD": c.password,
	}
	if c.clientSecret != "" {
		params["SECRET_HASH"] = c.secretHash(c.username)
	}

	out, err := c.client.InitiateAuth(context.Background(), &cip.InitiateAuthInput{
		AuthFlow:       ciptypes.AuthFlowTypeUserPasswordAuth,
		ClientId:       aws.String(c.clientID),
		AuthParameters: params,
	})
	if err != nil {
		return nil, fmt.Errorf("cognito: USER_PASSWORD_AUTH failed: %w", err)
	}
	if out.AuthenticationResult == nil {
		return nil, fmt.Errorf("cognito: USER_PASSWORD_AUTH returned nil result")
	}

	return c.storeResult(out.AuthenticationResult), nil
}

func (c *CognitoTokenSource) doRefresh() (*oauth2.Token, error) {
	params := map[string]string{
		"REFRESH_TOKEN": c.refreshToken,
	}
	if c.clientSecret != "" {
		params["SECRET_HASH"] = c.secretHash(c.username)
	}

	out, err := c.client.InitiateAuth(context.Background(), &cip.InitiateAuthInput{
		AuthFlow:       ciptypes.AuthFlowTypeRefreshTokenAuth,
		ClientId:       aws.String(c.clientID),
		AuthParameters: params,
	})
	if err != nil {
		return nil, fmt.Errorf("cognito: REFRESH_TOKEN_AUTH failed: %w", err)
	}
	if out.AuthenticationResult == nil {
		return nil, fmt.Errorf("cognito: REFRESH_TOKEN_AUTH returned nil result")
	}

	return c.storeResult(out.AuthenticationResult), nil
}

func (c *CognitoTokenSource) storeResult(res *ciptypes.AuthenticationResultType) *oauth2.Token {
	c.accessToken = *res.AccessToken
	c.expiry = time.Now().Add(time.Duration(res.ExpiresIn) * time.Second)
	if res.RefreshToken != nil {
		c.refreshToken = *res.RefreshToken
	}
	return &oauth2.Token{
		AccessToken: c.accessToken,
		Expiry:      c.expiry,
	}
}

// secretHash computes the HMAC-SHA256 secret hash required by Cognito when
// a client secret is configured.
func (c *CognitoTokenSource) secretHash(username string) string {
	mac := hmac.New(sha256.New, []byte(c.clientSecret))
	mac.Write([]byte(username + c.clientID))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
