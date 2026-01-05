package config

type GoogleCloudCredential struct{}

type GoogleCloud struct {
	OrganizationID string                `json:"organization_id" omitempty:"true"`
	ProjectID      string                `json:"project_id" omitempty:"true"`
	Firestore      Firestore             `json:"firestore" omitempty:"true"`
	PubSub         PubSub                `json:"pubsub" omitempty:"true"`
	Credential     GoogleCloudCredential `json:"credentials" omitempty:"true"`
}

type PubSub struct {
	Topics map[string]string `json:"topics" omitempty:"true"`
}

type Firestore struct {
	Database string `json:"database" omitempty:"true"`
}
