package config

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/api/idtoken"
)

type Crowemi struct {
	ClientName      string            `json:"client_name" omitempty:"true"`
	ClientID        string            `json:"client_id" omitempty:"true"`
	ClientSecretKey string            `json:"client_secret_key" omitempty:"true"`
	Uri             map[string]string `json:"uri" omitempty:"true"`
	DatabaseURI     string            `json:"database_uri" omitempty:"true"`
	Env             string            `json:"env" omitempty:"true"`
	Debug           bool              `json:"debug" omitempty:"true"`
}

func (c *Crowemi) CreateHeaders(req *http.Request, audience string, sessionID string) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("crowemi-client-id", c.ClientID)
	req.Header.Set("crowemi-client-secret-key", c.ClientSecretKey)
	req.Header.Set("crowemi-client-name", c.ClientName)
	req.Header.Set("crowemi-session-id", sessionID)
	if c.Env == "dev" || c.Env == "prod" {
		token, err := c.GetAuth(audience)
		if err != nil {
			fmt.Printf("Error getting auth token: %v\n", err)
			return err
		}
		req.Header.Set("Authorization", "Bearer "+token)
	}
	return nil
}

func (c *Crowemi) GetAuth(targetAudience string) (string, error) {
	ctx := context.Background()

	tokenSource, err := idtoken.NewTokenSource(ctx, targetAudience)
	if err != nil {
		fmt.Printf("Error creating token source: %v\n", err)
		return "", err
	}

	token, err := tokenSource.Token()
	if err != nil {
		fmt.Printf("Error retrieving identity token: %v\n", err)
		return "", err
	}

	return token.AccessToken, nil
}
