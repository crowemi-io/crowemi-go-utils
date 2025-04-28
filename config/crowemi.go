package config

type Crowemi struct {
	ClientName      string            `json:"client_name"`
	ClientID        string            `json:"client_id"`
	ClientSecretKey string            `json:"client_secret_key"`
	Uri             map[string]string `json:"uri"`
	DatabaseURI     string            `json:"database_uri"`
}
