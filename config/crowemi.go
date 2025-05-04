package config

type Crowemi struct {
	ClientName      string            `json:"client_name" omitempty:"true"`
	ClientID        string            `json:"client_id" omitempty:"true"`
	ClientSecretKey string            `json:"client_secret_key" omitempty:"true"`
	Uri             map[string]string `json:"uri" omitempty:"true"`
	DatabaseURI     string            `json:"database_uri" omitempty:"true"`
}
