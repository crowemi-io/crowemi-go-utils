package config

type GoogleCloudCredential struct{}

type GoogleCloud struct {
	ProjectId  string                `json:"project_id"`
	Topic      string                `json:"topic"`
	Credential GoogleCloudCredential `json:"credentials"`
}
