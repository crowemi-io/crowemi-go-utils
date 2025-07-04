package config

type Alpaca struct {
	AccountID    string `json:"account_id" omitempty:"true"`
	APIKey       string `json:"api_key" omitempty:"true"`
	APISecretKey string `json:"api_secret_key" omitempty:"true"`
	APIBaseURL   string `json:"api_base_url" omitempty:"true"`
	APIDataURL   string `json:"api_data_url" omitempty:"true"`
}
