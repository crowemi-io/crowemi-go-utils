package config

type GoogleCloudCredential struct{}

type GoogleCloud struct {
	ProjectId  string                `json:"project_id" omitempty:"true"`
	Topic      string                `json:"topic" omitempty:"true"`
	Credential GoogleCloudCredential `json:"credentials" omitempty:"true"`
}
