package config

type GoogleCloudCredential struct{}

type GoogleCloud struct {
	OrganizationID string                `json:"organization_id" omitempty:"true"`
	ProjectID      string                `json:"project_id" omitempty:"true"`
	Topics         map[string]string     `json:"topics" omitempty:"true"`
	Credential     GoogleCloudCredential `json:"credentials" omitempty:"true"`
}
