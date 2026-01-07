package config

type GoogleCloud struct {
	OrganizationID string      `json:"organization_id" omitempty:"true"`
	ProjectID      string      `json:"project_id" omitempty:"true"`
	Region         string      `json:"region" omitempty:"true"`
	Firestore      Firestore   `json:"firestore" omitempty:"true"`
	PubSub         PubSub      `json:"pubsub" omitempty:"true"`
	GenAI          GenAI       `json:"genai" omitempty:"true"`
	Credential     Credentials `json:"credentials" omitempty:"true"`
}

type Credentials struct{}
type GenAI struct {
	Model  string `json:"model" omitempty:"true"`
	ApiKey string `json:"api_key" omitempty:"true"`
}

type PubSub struct {
	Topics map[string]string `json:"topics" omitempty:"true"`
}

type Firestore struct {
	Database string `json:"database" omitempty:"true"`
}
