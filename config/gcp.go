package config

type GoogleCloudCredential struct{}

type GoogleCloud struct {
	ProjectID  string                `json:"project_id" omitempty:"true"`
	Topics     map[string]string     `json:"topics" omitempty:"true"`
	Credential GoogleCloudCredential `json:"credentials" omitempty:"true"`
}
